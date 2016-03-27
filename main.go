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

var (
	save    *bool
	saveDev *bool
)

func main() {

	save := flag.Bool("save", false, "Save to deps")
	saveDev := flag.Bool("save-dev", false, "Save to dev deps")

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	flag.Parse()

	fmt.Println(*save)
	fmt.Println(*saveDev)

	ParseSubCommands(os.Args, *save, *saveDev)
}
