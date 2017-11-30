package main

import (
	"flag"
	"log"
	"strings"

	"github.com/mccurdyc/websocket-example/client"
	"github.com/mccurdyc/websocket-example/server"
)

var (
	serverHost = flag.String("host", "localhost", "server host")
	serverPort = flag.Int("port", 8080, "server port")
	connType   = flag.String("connection", "server", "connection type [server|client]")
)

func init() {
	flag.Parse()
}

func main() {
	switch strings.ToLower(*connType) {
	case "client":
		log.Printf("client connecting to server running on %d\n", *serverPort)
		c := client.New()
		c.Connect(*serverHost, *serverPort)
		log.Printf("client connected to server running on %d\n", *serverPort)
	default:
		s := server.NewService(*serverHost, *serverPort)
		log.Printf("started server on %d\n", *serverPort)
		s.Start()
	}
}
