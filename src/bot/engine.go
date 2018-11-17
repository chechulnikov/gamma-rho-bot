package bot

type engine struct {
	listener       *listener
	spellChecker   *spellChecker
	commandHandler *commandHandler
}

func (e *engine) Start() {
	message := make(chan *chatMessage)

	go e.listener.start(message)

	for {
		go e.handleMessage(<-message)
	}
}

func (e *engine) handleMessage(message *chatMessage) {
	if message.command != nil {
		e.commandHandler.Handle(message.command)
		return
	}

	e.spellChecker.check(message)
}
