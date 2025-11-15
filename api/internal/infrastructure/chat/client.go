package chat

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	hub   *Hub
	conn  *websocket.Conn
	send  chan []byte
	chats map[int64]bool
}

type IncomingWS struct {
	Type   string `json:"type"` // "subscribe", "unsubscribe", "message", "history_request"
	ChatID int64  `json:"chat_id,omitempty"`
	Text   string `json:"text,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(1024 * 10)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("ws read error:", err)
			break
		}

		var in IncomingWS
		if err := json.Unmarshal(message, &in); err != nil {
			log.Println("bad ws json:", err)
			continue
		}

		switch in.Type {
		case "subscribe":
			if in.ChatID != 0 {
				c.hub.Subscribe(c, in.ChatID)
			}

		case "unsubscribe":
			if in.ChatID != 0 {
				c.hub.Unsubscribe(c, in.ChatID)
			}

		case "message":
			if in.ChatID == 0 || in.Text == "" {
				continue
			}
			if err := c.hub.HandleIncomingMessage(in.ChatID, in.Text); err != nil {
				log.Println("HandleIncomingMessage:", err)
			}

		case "history_request":
			if in.ChatID == 0 {
				continue
			}
			if in.Limit <= 0 {
				in.Limit = 50
			}
			if err := c.hub.SendHistory(c, in.ChatID, in.Limit); err != nil {
				log.Println("SendHistory:", err)
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("ws write error:", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}
