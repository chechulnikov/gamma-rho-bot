package spell

import (
	"fmt"
	"grammar-bot/telegram"
)

type sender struct {
	telegramClient telegram.BotAPIClient
	error          chan error
}

func (s *sender) send(message chatMessage) {
	err := s.telegramClient.SendMessageAsReply(message.chatId, message.text, message.replyToMessageId)
	if err != nil {
		s.error <- fmt.Errorf("sending message failed: %s", err.Error())
	}
}
