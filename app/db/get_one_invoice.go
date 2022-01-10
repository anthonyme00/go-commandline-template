package db

import (
	"cltest/commands"
	"cltest/models"
	"cltest/repository"
	"cltest/utils"
	"flag"
	"fmt"
)

type OneInvoiceCommandConfig struct {
	InvoiceId int
}

func (c *OneInvoiceCommandConfig) LoadConfig() {
	flag.IntVar(&c.InvoiceId, "invoiceid", 0, "invoiceid")
}

type OneInvoiceCommand struct {
	config OneInvoiceCommandConfig
}

func (cmd *OneInvoiceCommand) GetCommandDefinition() commands.CommandDefinition {
	return commands.CommandDefinition{
		Name:        "getinvoice",
		Description: "Get an invoice info by id",
		Usage:       "--invoiceid [invoiceid:int]",
	}
}

func (cmd *OneInvoiceCommand) Init() {
	cmd.config = OneInvoiceCommandConfig{}
	cmd.config.LoadConfig()
	flag.Parse()
}

func (cmd *OneInvoiceCommand) Execute(utils *utils.Utility, repo *repository.Repository) {
	invoice := models.Invoice{}
	err := invoice.GetInvoiceById(*repo.SQLRepository, cmd.config.InvoiceId)
	if err != nil {
		utils.Logger.LogError(err)
	} else {
		fmt.Println("Invoice Id : " + fmt.Sprintf("%d", invoice.InvoiceId))
		fmt.Println("Invoice No : " + invoice.InvoiceNo)
	}
}
