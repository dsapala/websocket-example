package client

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Client struct {
	Id   string
	Conn *websocket.Conn
}

func New() *Client {
	return &Client{
		Id: uuid.New().String(),
	}
}

func (c *Client) Connect(host, port string) error {
	dialer := &websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	var err error
	c.Conn, _, err = dialer.Dial(fmt.Sprintf("ws://%s%s/chat", host, port), nil)

	if err != nil {
		return err
	}

	c.writeMessages()

	var wg sync.WaitGroup

	// add this to a waitgroup because otherwise it returns before the goroutine returns
	wg.Add(1)
	go c.readMessages()
	wg.Wait()

	return nil
}

func (c *Client) writeMessages() {
	if err := c.Conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		log.Println(errors.Wrap(err, "error writing message to server"))
		return
	}
	return
}

func (c *Client) readMessages() {
	for {
		_, msg, err := c.Conn.ReadMessage()

		fmt.Printf("received: %v\n", string(msg))

		if err != nil {
			fmt.Println(errors.Wrap(err, "error reading message from server"))
			return
		}
	}
}
