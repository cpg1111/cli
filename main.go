package main

import (
	"flag"
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

const (
	libertyFile = "liberty.json"
	rootDir     = "_libs"
)

func main() {
	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args, flags)
}
