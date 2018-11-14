package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const telegramApiUri = "https://api.telegram.org/bot%s/%s"

type client struct {
	token string
}

func (c *client) GetUpdates(offset int64, limit int, timeout int, allowedUpdates []string) ([]Update, error) {
	uri := c.getMethodUrl("getUpdates")

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(struct {
		Offset         int64    `json:"offset"`
		Limit          int      `json:"limit"`
		Timeout        int      `json:"timeout"`
		AllowedUpdates []string `json:"allowed_updates"`
	}{Offset: offset, Limit: limit, Timeout: timeout, AllowedUpdates: allowedUpdates})

	response, err := http.Post(uri, contentTypeJSON, buffer)
	if err != nil {
		return nil, fmt.Errorf("Telegram API getUpdates method error: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Telegram API getUpdates method returns %d", response.StatusCode)
	}

	result := &updateResponse{}
	if err = json.NewDecoder(response.Body).Decode(result); err != nil {
		s := err.Error()
		return nil, fmt.Errorf("cann't decode JSON from response of Telegram API getUpdates method: %s", s)
	}

	return result.Result, nil
}

func (c *client) SendMessageAsReply(chatId int64, text string, replyToMessageId int64) error {
	uri := c.getMethodUrl("sendMessage")

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(struct {
		ChatId           int64  `json:"chat_id"`
		Text             string `json:"text"`
		ReplyToMessageId int64  `json:"reply_to_message_id"`
		ParseMode        string `json:"parse_mode"`
	}{
		ChatId:           chatId,
		Text:             text,
		ReplyToMessageId: replyToMessageId,
		ParseMode:        "Markdown",
	})

	if _, err := http.Post(uri, contentTypeJSON, buffer); err != nil {
		return fmt.Errorf("Telegram API sendMessage method error: %s", err)
	}

	return nil
}

func (c *client) getMethodUrl(methodName string) string {
	return fmt.Sprintf(telegramApiUri, c.token, methodName)
}

const (
	contentTypeJSON string = "application/json"
)

type updateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}
