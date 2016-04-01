package main

import (
	"os"
	"strings"

	"github.com/liberty-org/cli/utils"
)

func ExecuteRemove(subCmdArgs []string) {
	if len(subCmdArgs) <= 2 {
		utils.PrintErrorThenExit("Insufficient args to liberty remove", 1)
	}

	pathArg := subCmdArgs[2]
	removePackage(pathArg)
}

func removePackage(packageString string) {
	_, packagePath := utils.MakeGitPath(packageString)
	splitPath := strings.SplitAfter(packagePath, "/")

	libertyData := utils.ReadLibertyData()

	for i, elem := range libertyData.Dependencies {
		if elem.Name == packageString {
			libertyData.Dependencies = append(libertyData.Dependencies[:i], libertyData.Dependencies[i+1:]...)
		}
	}

	libertyData.GenerateLibertyFile()

	os.RemoveAll(utils.LibsDir + splitPath[0])
}
