package commands

import "fmt"

type CommandDefinition struct {
	Name        string
	Description string
	Group       string
	Usage       string
}

func (cd *CommandDefinition) UsageString() string {
	return fmt.Sprintf("%s %s %s", cd.Group, cd.Name, cd.Usage)
}

type Command interface {
	GetCommandDefinition() *CommandDefinition
	Init()
	Execute()
}
