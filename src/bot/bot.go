package bot

import (
	"fmt"
	"gamma-rho-bot/bing"
	"gamma-rho-bot/bot/command"
	"gamma-rho-bot/telegram"
)

type Settings struct {
	TelegramToken   string
	ChatsIds        map[int64]struct{}
	BingSpellAPIKey string
	Error           chan error
}

func New(settings Settings) (Engine, error) {
	telegramClient, err := telegram.NewBotAPIClient(settings.TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("can't construct Telegram Bot API client error: %s", err.Error())
	}

	spellCheckerAPIClient, err := bing.NewSpellCheckAPIClient(settings.BingSpellAPIKey)
	if err != nil {
		return nil, fmt.Errorf("can't construct Bing Spell Engine API client error: %s", err.Error())
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
	spellChecker := spellChecker{
		corrector: &corrector,
		sender:    &sender,
	}
	ivCommandExecutor := command.NewIVCommandExecutor()
	commandHandler := commandHandler{
		ivCommandExecutor: ivCommandExecutor,
		sender:            &sender,
		error:             settings.Error,
	}

	return &engine{
		listener:       &listener,
		spellChecker:   &spellChecker,
		commandHandler: &commandHandler,
	}, nil
}

type Engine interface {
	Start()
}
