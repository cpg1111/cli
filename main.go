package main

import (
	"fmt"
	"os"
)

const (
	libertyFile = "liberty.json"
	rootDir = "_libs"
)

func main() {

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
}
