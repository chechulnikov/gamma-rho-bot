package bot

import (
	"fmt"
	"gamma-rho-bot/telegram"
	"strings"
	"time"
)

type listener struct {
	telegramClient telegram.BotAPIClient
	updatesOffset  int64
	chatsIds       map[int64]struct{}
	error          chan error
}

func (l *listener) start(message chan *chatMessage) {
	for {
		updates, err := l.getUpdates()
		if err != nil {
			l.error <- fmt.Errorf("can't get bot updates from Telegram: %s", err.Error())
			time.Sleep(time.Second)
			continue
		}

		if len(updates) == 0 {
			continue
		}

		l.recalculateOffset(updates)

		for _, update := range updates {
			if msg, ok := l.tryToFormMessage(update); ok {
				message <- msg
			}
		}
	}
}

func (l *listener) getUpdates() ([]telegram.Update, error) {
	return l.telegramClient.GetUpdates(
		l.updatesOffset,
		100,
		60,
		[]string{"messages"},
	)
}

func (l *listener) recalculateOffset(updates []telegram.Update) {
	l.updatesOffset = updates[len(updates)-1].Id + 1
}

func (l *listener) tryToFormMessage(update telegram.Update) (*chatMessage, bool) {
	if l.chatsIds != nil {
		if _, ok := l.chatsIds[update.Message.Chat.Id]; !ok {
			return nil, false
		}
	}

	update.Message.Text = strings.TrimSpace(update.Message.Text)
	if update.Message.Text == "" {
		return nil, false
	}

	message := &chatMessage{
		id:     update.Message.Id,
		chatId: update.Message.Chat.Id,
		text:   update.Message.Text,
	}

	if botCommand, ok := tryToGetBotCommand(update); ok {
		message.command = botCommand
	}

	return message, true
}

func tryToGetBotCommand(update telegram.Update) (*botCommand, bool) {
	if len(update.Message.Entities) == 0 {
		return nil, false
	}

	if update.Message.Entities[0].Type != telegram.BotCommandMessageEntity {
		return nil, false
	}

	commandEntity := update.Message.Entities[0]

	command := update.Message.Text[commandEntity.Offset:len(update.Message.Text)]

	splittedCommand := strings.SplitN(command, " ", 2)

	if len(splittedCommand) < 1 {
		return nil, false
	}

	result := &botCommand{
		chatId: update.Message.Chat.Id,
		name:   strings.TrimPrefix(splittedCommand[0], "/"),
	}
	if len(splittedCommand) >= 2 {
		result.value = strings.TrimSpace(splittedCommand[1])
	}

	return result, true
}