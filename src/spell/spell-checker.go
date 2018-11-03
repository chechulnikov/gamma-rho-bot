package spell

import (
	"fmt"
	"time"
)

type spellChecker struct {
	listener  *listener
	corrector *corrector
	sender    *sender
}

func (sc *spellChecker) Start() {
	message := make(chan chatMessage)

	go sc.listener.start(message)

	for {
		receivedMessage := <-message
		isRevised, revisedText := sc.corrector.checkAndCorrect(receivedMessage.text)

		if isRevised {
			revisedText = fmt.Sprintf("✨ %s", revisedText)
			sc.sender.send(chatMessage{
				chatId:           receivedMessage.chatId,
				text:             revisedText,
				replyToMessageId: receivedMessage.id,
			})
		}
		time.Sleep(time.Second)
	}
}
