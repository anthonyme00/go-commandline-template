package main

import (
	"cltest/app/db"
	"cltest/app/test"
	"cltest/commands"
	"cltest/configs"
)

func main() {
	configs.Init()
	chain := commands.NewChain()

	test.Register(&chain)
	db.Register(&chain)

	chain.Execute()
}
