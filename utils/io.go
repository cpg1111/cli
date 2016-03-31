package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	LibertyFile = "liberty.json"
)

type LibertyData struct {
	Title        string
	Description  string
	Author       string
	Dependencies []Dependency
}

type Dependency struct {
	Name    string
	Version string
}

func pathInCwd(relativePath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cwd + "/" + relativePath
}

func ReadRepoDefinitions() []Repository {
	repoDefFile, err := ioutil.ReadFile(pathInCwd("repo_defs.json")) //Change to be a global file
	if err != nil {
		fmt.Println("Could not find valid repository definitions")
		os.Exit(2)
	}

	var repoDef []Repository
	json.Unmarshal(repoDefFile, &repoDef)

	return repoDef
}

func ReadLibertyData() LibertyData {
	libFile, err := ioutil.ReadFile(pathInCwd("liberty.json"))
	if err != nil {
		fmt.Println("Could not find valid liberty file")
		os.Exit(2)
	}

	var libertyData LibertyData
	json.Unmarshal(libFile, &libertyData)

	return libertyData
}

// Generates the Liberty File with the given data
// It saves it as "liberty.json" in the directory that the program was executed
func (libertyData *LibertyData) GenerateLibertyFile() {
	jsonData, marshalErr := json.MarshalIndent(libertyData, "", "\t")
	if marshalErr != nil {
		panic(marshalErr)
	}

	if err := ioutil.WriteFile(LibertyFile, jsonData, os.ModePerm); err != nil {
		panic(err)
	}
}

func WritePackageVersion(packageName string, version string) {
	libertyData := ReadLibertyData()

	for i, dep := range libertyData.Dependencies {
		if dep.Name == packageName {
			libertyData.Dependencies[i].Version = version
		}
	}

	libertyData.GenerateLibertyFile()
}
