package main

import (
	"os"

	"github.com/liberty-org/cli/utils"
)

const (
	getCmd     = "get"
	initCmd    = "init"
	listCmd    = "list"
	installCmd = "install"
	removeCmd  = "remove"
	helpCmd    = "help"
	setCmd     = "set"
)

func SubCommandsExist() bool {
	if len(os.Args) <= 1 {
		return false
	}
	return true
}

func ParseSubCommands(subCmdArgs []string) {
	switch subCmdArgs[1] {
	case helpCmd:
		helpCommand(subCmdArgs)
	case getCmd:
		getCommand(subCmdArgs)
	case initCmd:
		initCommand(subCmdArgs)
	case listCmd:
		listCommand(subCmdArgs)
	case installCmd:
		installCommand(subCmdArgs)
	case removeCmd:
		removeCommand(subCmdArgs)
	case setCmd:
		setCommand(subCmdArgs)
	default:
		utils.PrintErrorThenExit("Liberty does not recognise this command", 1)
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

func removeCommand(subCmdArgs []string) {
	ExecuteRemove(subCmdArgs)
}

func helpCommand(subCmdArgs []string) {
	ExecuteHelp(subCmdArgs)
}

func setCommand(subCmdArgs []string) {
	ExecuteSet(subCmdArgs)
}
