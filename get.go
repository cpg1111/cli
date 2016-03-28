package main

import (
	"fmt"
	"os"

	"github.com/liberty-org/cli/utils"
)

// ExecuteGet fetches a package from a remote source and places it in the _libs
// folder.
func ExecuteGet(args []string) string {
	if len(args) <= 2 {
		fmt.Println("Insufficient args to liberty get")
		os.Exit(2)
	}

	repoArg := args[2]
	gitRepo, userProjDir := utils.MakeGitPath(repoArg)
	localRepo := utils.CloneGitRepo(gitRepo, userProjDir)

	libertyData := ReadLibertyData()
	newDep := &Dependency{Name: repoArg}
	libertyData.Dependencies = append(libertyData.Dependencies, *newDep)

	libertyData.GenerateLibertyFile()

	return localRepo
}
