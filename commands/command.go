package commands

import (
	"cltest/repository"
	"cltest/utils"
)

type CommandDefinition struct {
	Name        string
	Description string
	Usage       string
}

type ICommand interface {
	GetCommandDefinition() CommandDefinition
	Init()
	Execute(util *utils.Utility, repo *repository.Repository)
}
