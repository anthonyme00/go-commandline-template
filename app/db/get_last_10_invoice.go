package db

import (
	"cltest/commands"
	"cltest/models"
	"cltest/repository"
	"cltest/utils"
	"flag"
	"fmt"
)

type Last10InvoiceCommandConfig struct {
}

func (c *Last10InvoiceCommandConfig) LoadConfig() {
}

type Last10InvoiceCommand struct {
	config Last10InvoiceCommandConfig
}

func (cmd *Last10InvoiceCommand) GetCommandDefinition() commands.CommandDefinition {
	return commands.CommandDefinition{
		Name:        "getlast10invoice",
		Description: "Get an invoice info by id",
		Usage:       "",
	}
}

func (cmd *Last10InvoiceCommand) Init() {
	cmd.config = Last10InvoiceCommandConfig{}
	cmd.config.LoadConfig()
	flag.Parse()
}

func (cmd *Last10InvoiceCommand) Execute(utils *utils.Utility, repo *repository.Repository) {
	invoices := models.InvoiceArray{}
	err := invoices.GetLast10Invoice(*repo.SQLRepository)
	if err != nil {
		utils.Logger.LogError(err)
	} else {
		for _, invoice := range invoices {
			fmt.Println("Invoice Id : " + fmt.Sprintf("%d", invoice.InvoiceId))
			fmt.Println("Invoice No : " + invoice.InvoiceNo)
			fmt.Println("----------------------------------")
			utils.Logger.LogMessage(fmt.Sprintf("Invoice : [%s, %d]", invoice.InvoiceNo, invoice.InvoiceId))
		}
	}
}
