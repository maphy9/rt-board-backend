package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	manager *Manager
	messageChannel chan []byte
}

type ClientList map[*Client]bool

func newClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn: conn,
		manager: manager,
		messageChannel: make(chan []byte, 256),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
		log.Println("Connection closed")
	}()

	for {
		_, payload, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		bytesPayload := []byte(payload)
		for client := range c.manager.clients {
			if client == c {
				continue
			}
			client.messageChannel <- bytesPayload
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
		log.Println("Connection closed")
	}()
		
	for message := range c.messageChannel {
		writer, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			break
		}
		writer.Write(message)

		if err := writer.Close(); err != nil {
			break
		}
	}
}
