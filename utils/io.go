package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	LibertyFile = "liberty.json"
	RepoDefFile = "repo_defs.json"
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
	repoDefFile, err := ioutil.ReadFile(pathInCwd(RepoDefFile)) //Change to be a global file
	if err != nil {
		PrintErrorThenExit("Could not find valid repository definitions", 2)
	}

	var repoDef []Repository
	json.Unmarshal(repoDefFile, &repoDef)

	return repoDef
}

func ReadLibertyData() LibertyData {
	libFile, err := ioutil.ReadFile(pathInCwd(LibertyFile))
	if err != nil {
		PrintErrorThenExit("There was a problem reading your Liberty file", 2)
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
		PrintErrorThenExit(marshalErr.Error(), 2)
	}

	if err := ioutil.WriteFile(LibertyFile, jsonData, os.ModePerm); err != nil {
		PrintErrorThenExit(err.Error(), 2)
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
