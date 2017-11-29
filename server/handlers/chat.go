package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type ConnectionPool struct {
	clients []*websocket.Conn

	// bi-directional channel
	channel chan []byte
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		clients: []*websocket.Conn{},
		channel: make(chan []byte),
	}
}

func (cp *ConnectionPool) add(c *websocket.Conn) {
	cp.clients = append(cp.clients, c)
	return
}

func (cp *ConnectionPool) readMessages() {
	// log the message to a file to create junk files

	for {
		for _, c := range cp.clients {
			_, msg, err := c.ReadMessage()

			if err != nil {
				log.Println(errors.Wrap(err, "error reading message from client"))
				return
			}

			fmt.Printf("message: %s\n", msg)
			// cp.channel <- string(msg)
		}
	}
}

// func (cp *ConnectionPool) writeMessages() {
// 	for {
// 		if err := cp.clientWriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

func Chat(w http.ResponseWriter, r *http.Request) {
	cp := NewConnectionPool()

	var u = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// upgrade to the WebSocket Protocol
	conn, err := u.Upgrade(w, r, nil)

	if err != nil {
		log.Println(errors.Wrap(err, "error upgrading to websocket connection"))
		return
	}

	// add upgraded connection to the connection pool
	cp.add(conn)

	// writeMessages(c)
	go cp.readMessages()
}
