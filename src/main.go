package main

import (
	"fmt"
	"gamma-rho-bot/bot"
	"log"
)

func main() {
	fmt.Println("γρ-bot started")

	config, e := loadConfig()
	if e != nil {
		panic(fmt.Errorf("can't load config: %s", e.Error()))
	}

	err := make(chan error)
	go logErrors(err)

	engine, e := bot.NewEngine(bot.Settings{
		TelegramToken:   config.telegramBotToken,
		ChatsIds:        config.chatsIds,
		BingSpellAPIKey: config.bingSpellAPIKey,
		Error:           err,
	})
	if e != nil {
		panic(fmt.Errorf("can't build bot: %s", e.Error()))
	}

	go engine.Start()

	<-make(chan struct{})
}

func logErrors(err chan error) {
	for {
		log.Println(<-err)
	}
}
