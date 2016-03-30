package main

import (
	"fmt"
	"os"

	"github.com/liberty-org/cli/utils"
)

func ExecuteSet(args []string) {
	if len(args) <= 3 {
		fmt.Println("Insufficient args to liberty setver")
		os.Exit(2)
	}

	packageName := args[2]
	packageVersion := args[3]
	repoPath := utils.MakePathFromPackage(packageName)
	utils.SetRepoVersion(repoPath, packageVersion)
}
