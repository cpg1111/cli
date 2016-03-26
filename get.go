package main

import (
	"fmt"
	"os"
)

// ExecuteGet fetches a package from a remote source and places it in the _libs
// folder.
func ExecuteGet(args []string) string {
	if len(args) <= 2 {
		fmt.Println("Insufficient args to liberty get")
		os.Exit(2)
	}

	repoArg := args[2]
	gitRepo, userProjDir := MakeGitPath(repoArg)
	localRepo := CloneGitRepo(gitRepo, userProjDir)

	return localRepo
}
