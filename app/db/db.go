package db

import "cltest/commands"

func Register(chain *commands.Chain) {
	grp := chain.Group("db", "Test DB Connections")
	grp.AddCommand(&OneInvoiceCommand{})
	grp.AddCommand(&Last10InvoiceCommand{})
}
