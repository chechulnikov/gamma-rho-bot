package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	envVarTelegramBotToken = "TELEGRAM_BOT_TOKEN"
	envVarBingSpellAPIKey  = "BING_SPELL_API_KEY"
	envVarChatsIdsCsv      = "CHATS_IDS_CSV"
)

type config struct {
	telegramBotToken string
	chatsIds         map[int64]struct{}
	bingSpellAPIKey  string
}

func loadConfig() (*config, error) {
	telegramBotToken := os.Getenv(envVarTelegramBotToken)
	if telegramBotToken == "" {
		return nil, fmt.Errorf("env var %s is empty", envVarTelegramBotToken)
	}

	bingSpellAPIKey := os.Getenv(envVarBingSpellAPIKey)
	if bingSpellAPIKey == "" {
		return nil, fmt.Errorf("env var %s is empty", envVarBingSpellAPIKey)
	}

	chatsIdsCsv := os.Getenv(envVarChatsIdsCsv)
	var chatsIds map[int64]struct{}
	var err error
	if chatsIdsCsv != "" {
		chatsIds, err = csvToHashSet(chatsIdsCsv)
		if err != nil {
			return nil, fmt.Errorf(
				"can't convert value of env var %s to map: %s",
				envVarChatsIdsCsv,
				err.Error(),
			)
		}
	}

	return &config{
		telegramBotToken: telegramBotToken,
		bingSpellAPIKey:  bingSpellAPIKey,
		chatsIds:         chatsIds,
	}, nil
}

func csvToHashSet(str string) (map[int64]struct{}, error) {
	result := make(map[int64]struct{})
	reader := csv.NewReader(strings.NewReader(str))
	for {
		values, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't read CSV: %s", err.Error())
		}

		for _, value := range values {
			integer, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"can't parse int from values %s: %s",
					values,
					err.Error(),
				)
			}
			result[integer] = struct{}{}
		}
	}

	return result, nil
}
