package main

import (
	"cltest/commands"
	"cltest/commands/test"
	"cltest/configs"
)

func main() {
	configs.Init()

	chain := commands.NewChain()

	test.Register(&chain)

	chain.Execute()
}
