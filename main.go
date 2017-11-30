package main

import (
	"flag"
	"log"
	"strings"

	"github.com/mccurdyc/websocket-example/client"
	"github.com/mccurdyc/websocket-example/server"
)

var (
	serverPort string
	serverHost string
	connType   string
)

func init() {
	flag.StringVar(&serverHost, "host", "localhost", "server host")
	flag.StringVar(&serverHost, "h", "localhost", "server host")
	flag.StringVar(&serverPort, "port", ":8080", "server port")
	flag.StringVar(&serverPort, "p", ":8080", "server port")
	flag.StringVar(&connType, "connection", "server", "connection type [server|client]")
	flag.StringVar(&connType, "c", "server", "connection type [server|client]")
	flag.Parse()
}

func main() {
	switch strings.ToLower(connType) {
	case "client":
		log.Printf("client connecting to server running on %s\n", serverPort)
		c := client.New()
		c.Connect(serverHost, serverPort)
		log.Printf("client connected to server running on %s\n", serverPort)
	default:
		s := server.NewService(serverPort)
		s.Start()
		log.Printf("started server on %s\n", serverPort)
	}
}
