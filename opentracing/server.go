package main

import (
	"context"
	"flag"
	"time"

	"github.com/kr/pretty"

	"github.com/halokid/rpcx-plus/serverplugin"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	// create a mock tracer
	tracer := mocktracer.New()
	opentracing.SetGlobalTracer(tracer)

	s := server.NewServer()
	p := serverplugin.OpenTracingPlugin{}
	s.Plugins.Add(p)
	s.RegisterName("Arith", new(Arith), "")

	go func() {
		time.Sleep(30 * time.Second)
		// print trace result
		spans := tracer.FinishedSpans()
		pretty.Print(spans)
	}()

	s.Serve("tcp", *addr)
}
