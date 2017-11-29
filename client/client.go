package client

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

	conn, _, err := dialer.Dial(fmt.Sprintf("ws://%s%s/chat", host, port), nil)

	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("hello there")); err != nil {
			return err
		}
	}

	return nil
}
