package test

import (
	"cltest/commands"
	"flag"
	"fmt"
)

type HelloWorldCommandConfig struct {
	Message string
	TestInt int
}

func (c *HelloWorldCommandConfig) LoadConfig() {
	flag.StringVar(&c.Message, "message", "Hello World!", "message")
	flag.IntVar(&c.TestInt, "testint", 0, "testint")
}

type HelloWorld struct {
	config HelloWorldCommandConfig
}

func (cmd *HelloWorld) GetCommandDefinition() commands.CommandDefinition {
	return commands.CommandDefinition{
		Name:        "helloworld",
		Description: "Print a message (defaults  to 'Hello World!')",
		Usage:       "--message [message:string] --testint [testint:int]",
	}
}

func (cmd *HelloWorld) Init() {
	cmd.config = HelloWorldCommandConfig{}
	cmd.config.LoadConfig()
	flag.Parse()
}

func (cmd *HelloWorld) Execute() {
	fmt.Println("Message : " + cmd.config.Message)
	fmt.Println("TestInt : " + fmt.Sprintf("%d", cmd.config.TestInt))
}
