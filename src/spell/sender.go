package spell

import (
	"fmt"
	"gamma-rho-bot/telegram"
	"log"
)

type sender struct {
	telegramClient telegram.BotAPIClient
	error          chan error
}

func (s *sender) send(message chatMessage) {
	log.Print("send message started...")
	err := s.telegramClient.SendMessageAsReply(message.chatId, message.text, message.replyToMessageId)
	if err != nil {
		s.error <- fmt.Errorf("sending message failed: %s", err.Error())
	}
	log.Print("send message finished")
}
