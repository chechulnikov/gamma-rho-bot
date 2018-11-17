package bot

type chatMessage struct {
	id               int64
	chatId           int64
	text             string
	replyToMessageId int64
	command          *botCommand
}

type botCommand struct {
	chatId int64
	name   string
	value  string
}
