package commands

import (
	"fmt"
	"os"
)

type commandGroup struct {
	name        string
	description string
}

type commandChain struct {
	commands map[string][]Command
}

func NewCommandChain() *commandChain {
	return &commandChain{
		commands: make(map[string][]Command),
	}
}

func (c *commandChain) AddCommand(cmd Command) *commandChain {
	if _, ok := c.commands[cmd.GetCommandDefinition().Group]; !ok {
		c.commands[cmd.GetCommandDefinition().Group] = make([]Command, 0)
	}
	c.commands[cmd.GetCommandDefinition().Group] = append(c.commands[cmd.GetCommandDefinition().Group], cmd)
	return c
}

func (c *commandChain) Execute() {
	printAvailableCommands := func() {
		for group, commands := range c.commands {
			fmt.Printf("group [%s]:\n", group)
			for _, command := range commands {
				fmt.Printf("\tcommand [%s] : \n\tDescription : %s\n\tUsage : %s\n", command.GetCommandDefinition().Name, command.GetCommandDefinition().Description, command.GetCommandDefinition().UsageString())
			}
		}
	}

	var cmd string
	var group string
	if len(os.Args) > 2 {
		group = os.Args[1]
		cmd = os.Args[2]
		os.Args = os.Args[2:]
	} else {
		fmt.Println("No command specified. Available commands:")
		printAvailableCommands()
		return
	}

	for grp, commands := range c.commands {
		if grp == group {
			for _, command := range commands {
				if command.GetCommandDefinition().Name == cmd {
					command.Init()
					command.Execute()
					return
				}
			}
		}
	}

	fmt.Printf("Command [%s] [%s] not found. Available commands:", group, cmd)
	printAvailableCommands()
}
