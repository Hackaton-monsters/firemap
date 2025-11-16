package response

import "time"

type CreatedMessage struct {
	Message Message `json:"message"`
}

type Message struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type Chats struct {
	Chats []Chat `json:"chats"`
}

type Chat struct {
	ID       int64     `json:"id"`
	Marker   Marker    `json:"marker"`
	Messages []Message `json:"messages"`
}
