package chat

import (
	"firemap/internal/domain/contract"
)

type subscription struct {
	client *Client
	chatID int64
}

type messageEvent struct {
	chatID int64
	data   []byte
}

type Hub struct {
	repository contract.MessageRepository

	register   chan *Client
	unregister chan *Client

	subscribe   chan subscription
	unsubscribe chan subscription

	broadcast  chan messageEvent
	chatEvents chan []byte

	clients map[*Client]bool
	rooms   map[int64]map[*Client]bool
}

func NewHub(repository contract.MessageRepository) *Hub {
	return &Hub{
		repository:  repository,
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		subscribe:   make(chan subscription),
		unsubscribe: make(chan subscription),
		broadcast:   make(chan messageEvent),
		chatEvents:  make(chan []byte),
		clients:     make(map[*Client]bool),
		rooms:       make(map[int64]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true

		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				for chatID, room := range h.rooms {
					if _, in := room[c]; in {
						delete(room, c)
						if len(room) == 0 {
							delete(h.rooms, chatID)
						}
					}
				}
				close(c.send)
			}

		case sub := <-h.subscribe:
			room := h.rooms[sub.chatID]
			if room == nil {
				room = make(map[*Client]bool)
				h.rooms[sub.chatID] = room
			}
			room[sub.client] = true
			if sub.client.chats == nil {
				sub.client.chats = make(map[int64]bool)
			}
			sub.client.chats[sub.chatID] = true

		case sub := <-h.unsubscribe:
			if room, ok := h.rooms[sub.chatID]; ok {
				if _, ok := room[sub.client]; ok {
					delete(room, sub.client)
					if len(room) == 0 {
						delete(h.rooms, sub.chatID)
					}
				}
			}
			delete(sub.client.chats, sub.chatID)

		case evt := <-h.broadcast:
			if room, ok := h.rooms[evt.chatID]; ok {
				for c := range room {
					select {
					case c.send <- evt.data:
					default:
						delete(room, c)
						delete(h.clients, c)
						close(c.send)
					}
				}
			}

		case data := <-h.chatEvents:
			for c := range h.clients {
				select {
				case c.send <- data:
				default:
					delete(h.clients, c)
					for chatID, room := range h.rooms {
						if _, in := room[c]; in {
							delete(room, c)
							if len(room) == 0 {
								delete(h.rooms, chatID)
							}
						}
					}
					close(c.send)
				}
			}
		}
	}
}

func (h *Hub) Subscribe(c *Client, chatID int64) {
	h.subscribe <- subscription{client: c, chatID: chatID}
}

func (h *Hub) Unsubscribe(c *Client, chatID int64) {
	h.unsubscribe <- subscription{client: c, chatID: chatID}
}

//func (h *Hub) HandleIncomingMessage(chatID int64, text string) error {
//	msg, err := h.repository.Add(...)
//	msg, err := h.repository.CreateMessage(context.Background(), chatID, nil, text)
//	if err != nil {
//		return err
//	}
//
//	payload, err := json.Marshal(struct {
//		Type    string   `json:"type"`
//		Payload *Message `json:"payload"`
//	}{
//		Type:    "message",
//		Payload: msg,
//	})
//	if err != nil {
//		return err
//	}
//
//	h.broadcast <- messageEvent{
//		chatID: chatID,
//		data:   payload,
//	}
//	return nil
//}
//
//// История по запросу клиента (по WS)
//func (h *Hub) SendHistory(c *Client, chatID int64, limit int) error {
//	msgs, err := h.repository.GetMessages(context.Background(), chatID, limit)
//	if err != nil {
//		return err
//	}
//
//	payload, err := json.Marshal(struct {
//		Type    string `json:"type"`
//		Payload struct {
//			ChatID   int64     `json:"chat_id"`
//			Messages []Message `json:"messages"`
//		} `json:"payload"`
//	}{
//		Type: "history",
//		Payload: struct {
//			ChatID   int64     `json:"chat_id"`
//			Messages []Message `json:"messages"`
//		}{
//			ChatID:   chatID,
//			Messages: msgs,
//		},
//	})
//	if err != nil {
//		return err
//	}
//
//	c.send <- payload
//	return nil
//}
//
//// Вызов из HTTP-хендлера, когда кто-то создал чат через API
//func (h *Hub) BroadcastChatCreated(chat *Chat) {
//	data, err := json.Marshal(struct {
//		Type    string `json:"type"`
//		Payload *Chat  `json:"payload"`
//	}{
//		Type:    "chat_created",
//		Payload: chat,
//	})
//	if err != nil {
//		log.Println("BroadcastChatCreated marshal:", err)
//		return
//	}
//	h.chatEvents <- data
//}
