package main

import (
	"fmt"
	"os"
)

type LibertyFile struct {
	Project         Project
	Dependencies    map[string]string
	DevDependencies map[string]string
}

type Project struct {
	Title       string
	Description string
	Author      string
}

func main() {
	fmt.Println("Hello, Liberty!")

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
}
