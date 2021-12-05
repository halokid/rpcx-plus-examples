package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	cancelFn()

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
