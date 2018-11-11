package spell

import (
	"fmt"
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
		go sc.handleMessage(<-message)
	}
}

func (sc *spellChecker) handleMessage(message chatMessage) {
	isRevised, revisedText := sc.corrector.checkAndCorrect(message.text)
	if isRevised {
		revisedText = fmt.Sprintf("✨ %s", revisedText)
		sc.sender.send(chatMessage{
			chatId:           message.chatId,
			text:             revisedText,
			replyToMessageId: message.id,
		})
	}
}
