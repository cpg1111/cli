package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/liberty-org/cli/utils"
)

func ExecuteInstall(subCmdArgs []string) {
	libertyData := ReadLibertyData()

	for _, dep := range libertyData.Dependencies {
		path, userProjDir := utils.MakeGitPath(dep.Name)
		utils.CloneGitRepo(path, userProjDir)
	}
}

func ReadLibertyData() LibertyData {
	libFile, err := ioutil.ReadFile("./liberty.json")
	if err != nil {
		fmt.Println("Could not find valid liberty file")
		os.Exit(2)
	}

	var libertyData LibertyData
	json.Unmarshal(libFile, &libertyData)

	return libertyData
}
