package main

import (
	"context"
	"flag"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server address")
	addr2 = flag.String("addr2", "baidu.com:80", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

type Arith2 int

func (t *Arith2) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	flag.Parse()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith), "weight=7")
		s.Serve("tcp", *addr1)
	}()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith2), "weight=3")
		s.Serve("tcp", *addr2)
	}()

	select {}
}
