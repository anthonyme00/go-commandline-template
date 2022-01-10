package main

import (
	"cltest/commands"
	"cltest/commands/test"
)

func main() {
	chain := commands.NewCommandChain()
	chain.AddCommand(&test.HelloWorld{}).Execute()
}

// var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// // This file was generated by robots at
// // {{ .Timestamp }}
// // using data from
// // {{ .URL }}
// package project

// var Contributors = []string{
// {{- range .Carls }}
// 	{{ printf "%q" . }},
// {{- end }}
// }
// `))