package command

func NewIVCommandExecutor() Executor {
	ivs := getIrregularVerbs()
	return &ivCommandExecutor{ivs: ivs}
}

type Selector interface {
	GetExecutor(command string) Executor
}

type Executor interface {
	Execute(value string) string
}
