package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/halokid/rpcx-plus/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	oneClient := client.NewOneClient(client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer oneClient.Close()

	f, err := os.Create("abc.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = oneClient.DownloadFile(context.Background(), "abc.txt", f)
	if err != nil {
		panic(err)
	}
	log.Println("received")

}
