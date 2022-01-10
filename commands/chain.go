package commands

import (
	"cltest/repository"
	"cltest/repository/sql"
	"cltest/utils"
	"fmt"
	"os"
)

type Chain struct {
	Groups []*Group
}

func NewChain() Chain {
	return Chain{
		Groups: make([]*Group, 0),
	}
}

func (c *Chain) Group(name, description string) *Group {
	grp := &Group{
		name:        name,
		description: description,
		commands:    make([]ICommand, 0),
	}
	c.Groups = append(c.Groups, grp)
	return grp
}

type Group struct {
	name        string
	description string
	commands    []ICommand
}

func (g *Group) AddCommand(command ICommand) {
	g.commands = append(g.commands, command)
}

func (c *Chain) Execute() {
	printAvailableCommands := func() {
		for _, group := range c.Groups {
			fmt.Printf("group [%s]:\n", group.name)
			fmt.Printf("%s\n  ", group.description)
			for _, command := range group.commands {
				usageString := fmt.Sprintf("%s %s %s", group.name, command.GetCommandDefinition().Name, command.GetCommandDefinition().Usage)
				fmt.Printf("\tcommand [%s] : \n\t  Description : %s\n\t  Usage : %s\n", command.GetCommandDefinition().Name, command.GetCommandDefinition().Description, usageString)
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

	for _, grp := range c.Groups {
		if grp.name == group {
			for _, command := range grp.commands {
				if command.GetCommandDefinition().Name == cmd {
					utility := utils.NewUtility(grp.name + "." + cmd)
					defer utility.Clean()

					sql, err := sql.NewMySQLRepository()
					if err != nil {
						utility.Logger.LogErrorMessage(err, "Failed to create MySQL repository")
						return
					}
					repo := repository.NewRepository(sql)
					defer repo.Clean()

					command.Init()
					command.Execute(utility, repo)
					return
				}
			}
		}
	}

	fmt.Printf("Command [%s] [%s] not found. Available commands:", group, cmd)
	printAvailableCommands()
}
