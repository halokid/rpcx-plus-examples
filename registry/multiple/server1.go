package main

// import (
// 	"context"
// 	"flag"

// 	example "github.com/halokid/rpcx-examples"
// 	"github.com/halokid/rpcx-plus/server"
// )

// var (
// 	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
// )

// type Arith2 int

// func (t *Arith2) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
// 	reply.C = args.A * args.B * 100
// 	return nil
// }
// func main() {
// 	flag.Parse()

// 	s := server.NewServer()
// 	s.RegisterName("Arith", new(example.Arith), "")
// 	s.Serve("tcp", *addr1)

// 	select {}
// }
