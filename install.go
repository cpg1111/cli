package main

import "github.com/liberty-org/cli/utils"

func ExecuteInstall(subCmdArgs []string) {
	libertyData := utils.ReadLibertyData()

	for _, dep := range libertyData.Dependencies {
		path, userProjDir := utils.MakeGitPath(dep.Name)
		utils.CloneGitRepo(path, userProjDir)
	}
}
