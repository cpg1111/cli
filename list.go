package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ExecuteList will print the entire tree of dependencies to the console
func ExecuteList(args []string) {
	dependencies := trackFolder(rootDir)
	fmt.Printf("%v", dependencies)
}

func trackFolder(path string) string {
	files, err := ioutil.ReadDir(path)
	fullPath := ""

	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(path + "/" + libertyFile); err == nil {
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
