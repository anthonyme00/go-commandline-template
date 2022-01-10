package commands

type CommandDefinition struct {
	Name        string
	Description string
	Usage       string
}

type Command interface {
	GetCommandDefinition() CommandDefinition
	Init()
	Execute()
}
