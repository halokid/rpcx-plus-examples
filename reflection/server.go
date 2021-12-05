package main

import (
	"flag"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/reflection"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := reflection.New()
	s.Plugins.Add(p)

	s.Register(new(example.Arith), "")
	s.Register(p, "")
	s.Serve("tcp", *addr)
}
