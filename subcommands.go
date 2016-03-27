package main

import (
	"fmt"
	"os"
)

const (
	getCmd     = "get"
	initCmd    = "init"
	listCmd    = "list"
	installCmd = "install"
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
	case listCmd:
		listCommand(subCmdArgs)
	case installCmd:
		installCommand(subCmdArgs)
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

func listCommand(subCmdArgs []string) {
	ExecuteList(subCmdArgs)
}

func installCommand(subCmdArgs []string) {
	ExecuteInstall(subCmdArgs)
}
