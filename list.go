package main

import (
	"io/ioutil"
	"os"

	"github.com/liberty-org/cli/utils"
)

// ExecuteList will print the entire tree of dependencies to the console
func ExecuteList(args []string) {
	dependencies := trackFolder(rootDir)
	utils.PrintInfo(dependencies)
}

func trackFolder(path string) string {
	files, err := ioutil.ReadDir(path)
	fullPath := ""

	if err != nil {
		utils.PrintUrgent("Could not find any dependencies")
		os.Exit(2)
	}

	if _, err := os.Stat(path + "/" + utils.LibertyFile); err == nil {
		return removeRootPrefix(path) + "\n"
	}

	for _, file := range files {
		if file.IsDir() {
			newPath := path + "/" + file.Name()
			fullPath += trackFolder(newPath)
		}
	}

	return fullPath
}

// Removes the `rootDir` prefix from the project names
func removeRootPrefix(path string) string {
	return path[len(rootDir)+1 : len(path)]
}
