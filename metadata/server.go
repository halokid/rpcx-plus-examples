package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/halokid/rpcx-plus/share"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)

	fmt.Printf("received meta: %+v\n", reqMeta)

	resMeta["echo"] = "from server"

	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}
