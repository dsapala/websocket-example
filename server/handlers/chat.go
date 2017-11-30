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
	broadcast chan []byte
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		clients:   []*websocket.Conn{},
		broadcast: make(chan []byte),
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
				// if we encounter errors, need to remove client from cp.clients
				log.Println(errors.Wrap(err, "error reading message from client"))
				return
			}

			cp.broadcast <- msg
		}
	}
}

func (cp *ConnectionPool) writeMessages() {
	msg := <-cp.broadcast

	for _, c := range cp.clients {
		fmt.Printf("message: %v\n", string(msg))
		if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
			// if we encounter errors, need to remove client from cp.clients
			log.Println(errors.Wrap(err, "error writing message to all clients"))
			return
		}
	}
}

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
	go cp.writeMessages()
}
