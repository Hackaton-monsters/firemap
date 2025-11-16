package response

import (
	"sort"
	"time"
)

type CreatedMessage struct {
	ChatID  int64   `json:"chatId"`
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
	ID        int64     `json:"id"`
	Marker    Marker    `json:"marker"`
	Messages  []Message `json:"messages"`
	CreatedAt time.Time `json:"created_at"`
}

func lastActivityTime(c Chat) time.Time {
	last := c.CreatedAt

	for i := range c.Messages {
		if c.Messages[i].CreatedAt.After(last) {
			last = c.Messages[i].CreatedAt
		}
	}

	return last
}

func (cs *Chats) SortByLastActivityDesc() {
	sort.Slice(cs.Chats, func(i, j int) bool {
		ti := lastActivityTime(cs.Chats[i])
		tj := lastActivityTime(cs.Chats[j])

		return ti.After(tj)
	})
}
