package test

import "cltest/commands"

func Register(chain *commands.Chain) {
	grp := chain.Group("test", "Test commands")
	grp.AddCommand(&HelloWorld{})
}
