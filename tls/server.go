package main

import (
	"crypto/tls"
	"flag"
	"log"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Print(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)
}
