package main

import (
	"fmt"
	"grammar-bot/spell"
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

	spellChecker, e := spell.NewChecker(spell.Settings{
		TelegramToken:   config.telegramBotToken,
		ChatsIds:        config.chatsIds,
		BingSpellAPIKey: config.bingSpellAPIKey,
		Error:           err,
	})
	if e != nil {
		panic(fmt.Errorf("can't construct spell checker spell: %s", e.Error()))
	}

	go spellChecker.Start()

	<-make(chan struct{})
}

func logErrors(err chan error) {
	e := <-err
	log.Println(e)
}
