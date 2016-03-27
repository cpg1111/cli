package main

import (
	"fmt"
	"os"
)

const (
	getCmd      = "get"
	initCmd     = "init"
	listCmd     = "list"
	saveFlag    = "--save"
	saveDevFlag = "--save-dev"
)

func SubCommandsExist() bool {
	if len(os.Args) <= 1 {
		return false
	}
	return true
}

func ParseSubCommands(subCmdArgs []string, save bool, saveDev bool) {
	switch subCmdArgs[1] {
	case getCmd:
		getCommand(subCmdArgs, save, saveDev)
	case initCmd:
		initCommand(subCmdArgs)
	case listCmd:
		listCommand(subCmdArgs)
	default:
		fmt.Println("Liberty does not recognise this command")
		os.Exit(2)
	}
}

func getCommand(subCmdArgs []string, save bool, saveDev bool) string {
	return ExecuteGet(subCmdArgs, save, saveDev)
}

func initCommand(subCmdArgs []string) {
	ExecuteInit(subCmdArgs)
}

func listCommand(subCmdArgs []string) {
	ExecuteList(subCmdArgs)
}
