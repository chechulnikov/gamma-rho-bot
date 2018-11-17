package bot

import (
	"fmt"
)

type spellChecker struct {
	corrector *corrector
	sender    *sender
}

func (sc *spellChecker) check(message *chatMessage) {
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
