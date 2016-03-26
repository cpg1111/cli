package main

import (
	"fmt"
	"os"
)

const (
	getCmd  = "get"
	initCmd = "init"
)

func SubCommandsExist() bool {
	if len(os.Args) <= 1 {
		return false
	}
	return true
}

func ParseSubCommands(subCmdArgs []string) {
	switch subCmdArgs[1] {
	case getCmd:
		getCommand(subCmdArgs)
	case initCmd:
		initCommand(subCmdArgs)
	default:
		fmt.Println("Liberty does not recognise this command")
		os.Exit(2)
	}
}

func getCommand(subCmdArgs []string) string {
	return ExecuteGet(subCmdArgs)
}

func initCommand(subCmdArgs []string) {
	ExecuteInit(subCmdArgs)
}
