package main

import (
	"cltest/commands"
	"cltest/commands/test"
)

func main() {
	chain := commands.NewChain()

	test.Register(&chain)

	chain.Execute()
}
