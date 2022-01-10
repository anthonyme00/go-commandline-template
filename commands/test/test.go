package test

import (
	"cltest/commands"
	"flag"
	"fmt"
)

type HelloWorldCommandConfig struct {
	Message string
}

func (c *HelloWorldCommandConfig) LoadConfig() {
	flag.StringVar(&c.Message, "message", "Hello World!", "message")
}

type HelloWorld struct {
	config HelloWorldCommandConfig
}

func (cmd *HelloWorld) GetCommandDefinition() *commands.CommandDefinition {
	return &commands.CommandDefinition{
		Name:        "helloworld",
		Description: "Print a message (defaults  to 'Hello World!')",
		Group:       "test",
		Usage:       "--message [message:string]",
	}
}

func (cmd *HelloWorld) Init() {
	cmd.config = HelloWorldCommandConfig{}
	cmd.config.LoadConfig()
	flag.Parse()
}

func (cmd *HelloWorld) Execute() {
	fmt.Println(cmd.config.Message)
}
