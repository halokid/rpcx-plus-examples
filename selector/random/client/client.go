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
	addr1 = flag.String("addr1", "tcp@localhost:8972", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8973", "server address")
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(time.Second)
	}

}
