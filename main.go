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
	cmd := os.Args[1]  // Program Name is always the first (implicit) argument

	switch cmd {
	case "help":
		PrintHelpArgs()
	}
	fmt.Printf("Program Name: %s\n", cmd)

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
}
