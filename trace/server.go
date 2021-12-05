package main

import (
	"flag"
	"net/http"

	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
	"github.com/halokid/rpcx-plus/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	go http.ListenAndServe(":8080", nil)

	s := server.NewServer()

	traceP := &serverplugin.OpenTracingPlugin{}
	//trace.AuthRequest = func(req *http.Request) (any, sensitive bool) { return true, true }

	s.Plugins.Add(traceP)

	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}
