package main

import (
	"fmt"
	"gamma-rho-bot/bot"
	"log"
)

func main() {
	fmt.Println("Grammar Bot started")

	config, e := loadConfig()
	if e != nil {
		panic(fmt.Errorf("can't load config: %s", e.Error()))
	}

	err := make(chan error)
	go logErrors(err)

	spellChecker, e := bot.New(bot.Settings{
		TelegramToken:   config.telegramBotToken,
		ChatsIds:        config.chatsIds,
		BingSpellAPIKey: config.bingSpellAPIKey,
		Error:           err,
	})
	if e != nil {
		panic(fmt.Errorf("can't construct bot checker bot: %s", e.Error()))
	}

	go spellChecker.Start()

	<-make(chan struct{})
}

func logErrors(err chan error) {
	e := <-err
	log.Println(e)
}
