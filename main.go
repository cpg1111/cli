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

	var save bool
	var saveDev bool

	flag.BoolVar(&save, "save", false, "Save to deps")
	flag.BoolVar(&saveDev, "save-dev", false, "Save to dev deps")

	flag.Parse()

	if !SubCommandsExist() {
		fmt.Println("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	fmt.Println(save)
	fmt.Println(saveDev)

	ParseSubCommands(os.Args, save, saveDev)
}
