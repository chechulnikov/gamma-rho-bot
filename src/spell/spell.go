package spell

import (
	"fmt"
	"gamma-rho-bot/bing"
	"gamma-rho-bot/telegram"
)

type Settings struct {
	TelegramToken   string
	ChatsIds        map[int64]struct{}
	BingSpellAPIKey string
	Error           chan error
}

func NewChecker(settings Settings) (Checker, error) {
	telegramClient, err := telegram.NewBotAPIClient(settings.TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("can't construct Telegram Bot API client error: %s", err.Error())
	}

	spellCheckerAPIClient, err := bing.NewSpellCheckAPIClient(settings.BingSpellAPIKey)
	if err != nil {
		return nil, fmt.Errorf("can't construct Bing Spell Checker API client error: %s", err.Error())
	}

	listener := listener{
		telegramClient: telegramClient,
		error:          settings.Error,
		chatsIds:       settings.ChatsIds,
	}
	corrector := corrector{
		spellCheckerAPIClient: spellCheckerAPIClient,
		error:                 settings.Error,
	}
	sender := sender{
		telegramClient: telegramClient,
		error:          settings.Error,
	}

	return &spellChecker{
		listener:  &listener,
		corrector: &corrector,
		sender:    &sender,
	}, nil
}

type Checker interface {
	Start()
}

type chatMessage struct {
	id               int64
	chatId           int64
	text             string
	replyToMessageId int64
}
