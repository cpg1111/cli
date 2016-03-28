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
