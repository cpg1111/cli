package main

import "github.com/liberty-org/cli/utils"

// ExecuteGet fetches a package from a remote source and places it in the _libs
// folder.
func ExecuteGet(args []string) string {
	if len(args) <= 2 {
		utils.PrintErrorThenExit("Insufficient args to liberty get", 2)
	}

	repoArg := args[2]
	gitRepo, userProjDir := utils.MakeGitPath(repoArg)

	utils.PrintInfof("Attempting to fetch remote dependency %v", gitRepo)

	localRepo := utils.CloneGitRepo(gitRepo, userProjDir)

	utils.PrintInfo("Updating your Liberty File")

	libertyData := utils.ReadLibertyData()
	newDep := &utils.Dependency{Name: repoArg}

	libertyData.Dependencies = append(libertyData.Dependencies, *newDep)
	libertyData.GenerateLibertyFile()

	utils.PrintSuccess("Operation Success. Package Installed")

	return localRepo
}
