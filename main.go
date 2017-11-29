package main

import (
	"flag"
	"log"
	"strings"

	"github.com/mccurdyc/websocket-example/client"
	"github.com/mccurdyc/websocket-example/server"
	"github.com/pkg/errors"
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
		c := client.New()
		if err := c.Connect(serverHost, serverPort); err != nil {
			log.Println(errors.Wrap(err, "error connecting to server"))
			return
		}
		log.Printf("client connected to server running on %s\n", serverPort)
	default:
		s := server.NewService(serverPort)
		s.Start()
		log.Printf("started server on %s\n", serverPort)
	}
}
