package command

type SendMessage struct {
	ChatID int64  `json:"chat_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
