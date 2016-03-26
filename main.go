package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, Liberty!")

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
}
