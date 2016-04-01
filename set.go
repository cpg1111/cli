package main

import "github.com/liberty-org/cli/utils"

func ExecuteSet(args []string) {
	if len(args) <= 3 {
		utils.PrintErrorThenExit("Insufficient args to liberty set", 1)
	}

	packageNameFull := args[2]
	packageName := utils.TrimProviderPrefix(args[2])
	packageVersion := args[3]
	repoPath := utils.MakePathFromPackage(packageName)
	utils.SetRepoVersion(repoPath, packageVersion)
	utils.WritePackageVersion(packageNameFull, packageVersion)
}
