package telegram

import "errors"

func NewBotAPIClient(token string) (BotAPIClient, error) {
	if token == "" {
		return nil, errors.New("token shouldn't be empty")
	}
	return &client{token: token}, nil
}

type BotAPIClient interface {
	GetUpdates(offset int64, limit int, timeout int, allowedUpdates []string) ([]Update, error)
	SendMessageAsReply(chatId int64, text string, replyToMessageId int64) error
}

type Update struct {
	Id      int64   `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Id   int64  `json:"message_id"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	Id int64 `json:"id"`
}
