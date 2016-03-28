package main

import (
	"fmt"
	"os"
	"strings"
)

func ExecuteRemove(subCmdArgs []string) {
	if len(subCmdArgs) <= 2 {
		fmt.Println("Insufficient args to liberty remove")
		os.Exit(2)
	}

	pathArg := subCmdArgs[2]
	_, userProjDir := MakeGitPath(pathArg)

	removePackage(userProjDir)
}

func removePackage(packagePath string) {
	splitPath := strings.SplitAfter(packagePath, "/")

	os.RemoveAll(libsDir + splitPath[0])
}
