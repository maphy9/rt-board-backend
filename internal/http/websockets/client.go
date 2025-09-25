package websockets

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn

	manager *Manager
}

type ClientList map[*Client]bool

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn: conn,
		manager: manager,
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
		log.Println("Connection closed")
	}()

	for {
		messageType, payload, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}
		log.Println("messageType: ", messageType)
		log.Println("Payload: ", string(payload))
	}
}

