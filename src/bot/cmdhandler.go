package bot

import (
	"fmt"
	"gamma-rho-bot/bot/command"
)

type commandHandler struct {
	ivCommandExecutor command.Executor
	sender            *sender
	error             chan error
}

func (h *commandHandler) Handle(command *botCommand) {
	executor := h.getExecutor(command.name)
	if executor == nil {
		return
	}

	response := executor.Execute(command.value)

	h.sender.send(chatMessage{
		chatId: command.chatId,
		text:   response,
	})
}

func (h *commandHandler) getExecutor(command string) command.Executor {
	switch command {
	case "iv":
		return h.ivCommandExecutor
	default:
		h.error <- fmt.Errorf("executor for command \"%s\" not found", command)
		return nil
	}
}
