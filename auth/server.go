package main

import (
	"context"
	"errors"
	"flag"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/protocol"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.AuthFunc = auth
	s.Serve("reuseport", *addr)
}

func auth(ctx context.Context, req *protocol.Message, token string) error {

	if token == "bearer tGzv3JOkF0XG5Qx2TlKWIA" {
		return nil
	}

	return errors.New("invalid token")
}
