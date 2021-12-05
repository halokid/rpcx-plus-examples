package main

import (
	"context"
	"flag"
	"fmt"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addrx = flag.String("addr", "localhost:8972", "server address")
)

type Arithx struct{}

// the second parameter is not a pointer
func (t *Arithx) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	fmt.Println("C=", reply.C)
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arithx", new(Arithx), "")
	err := s.Serve("tcp", *addrx)
	//err := s.Serve("http", *addrx)
	if err != nil {
		panic(err)
	}
}
