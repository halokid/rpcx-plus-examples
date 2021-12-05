package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"flag"
	"log"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/client"
	"github.com/halokid/rpcx-plus/protocol"
	"github.com/halokid/rpcx-plus/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	// register customized codec
	share.Codecs[protocol.SerializeType(4)] = &GobCodec{}
	option := client.DefaultOption
	option.SerializeType = protocol.SerializeType(4)

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}

type GobCodec struct {
}

func (c *GobCodec) Decode(data []byte, i interface{}) error {
	enc := gob.NewDecoder(bytes.NewBuffer(data))
	err := enc.Decode(i)
	return err
}

func (c *GobCodec) Encode(i interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(i)
	return buf.Bytes(), err
}
