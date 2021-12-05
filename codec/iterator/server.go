package main

import (
	"flag"

	"github.com/halokid/rpcx-plus/protocol"
	"github.com/halokid/rpcx-plus/share"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-examples/codec/iterator/codec"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	share.Codecs[protocol.SerializeType(4)] = &codec.JsoniterCodec{}
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}
