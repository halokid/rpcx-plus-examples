package main

import (
	"flag"
	"net"
	"time"

	graphite "github.com/cyberdelia/go-metrics-graphite"
	metrics "github.com/rcrowley/go-metrics"
	example "github.com/halokid/rpcx-examples"
	"github.com/halokid/rpcx-plus/server"
	"github.com/halokid/rpcx-plus/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2003")
	go graphite.Graphite(metrics.DefaultRegistry, 1e9, "rpcx.services.host.127_0_0_1", addr)
}
