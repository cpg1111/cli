package main

import (
	"fmt"
	"os"
)

type LibertyFile struct {
	Project         Project
	Dependencies    []Dependency
	DevDependencies []Dependency
}

type Project struct {
	Title       string
	Description string
	Author      string
}

type Dependency struct {
	Name    string
	version string
}

func main() {
	fmt.Println("Hello, Liberty!")

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
}
