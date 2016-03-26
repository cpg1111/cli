package main

import (
	"fmt"
	"os"
)

const (
	getCmd = "get"
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
	default:
		fmt.Println("Liberty does not recognise this command")
		os.Exit(2)
	}
}

func getCommand(subCmdArgs []string) string {
	if len(subCmdArgs) <= 2 {
		fmt.Println("Insufficient args to liberty get")
		os.Exit(2)
	}

	repoArg := subCmdArgs[2]
	gitRepo, userProjDir := MakeGitPath(repoArg)
	localRepo := CloneGitRepo(gitRepo, userProjDir)

	return localRepo
}
