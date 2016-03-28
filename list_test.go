package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
)

func TestThatListWorksForValidTree(t *testing.T) {
	defer cleanupListTest()

	firstPackage := "testauthor/testpackage"
	secondPackage := "testauthortwo/testpackagetwo"

	if err := createMockPackage(firstPackage, true); err != nil {
		t.Error(err)
	}

	if err := createMockPackage(secondPackage, true); err != nil {
		t.Error(err)
	}

	directoryList := trackFolder(rootDir)

	if directoryList != firstPackage+"\n"+secondPackage+"\n" {
		t.Error("The list command should print out the directories under _libs that have a liberty.json")
	}
}

// Creates a mock package in the _libs folder and also adds a liberty file
func createMockPackage(packageAlias string, addLibertyFile bool) error {
	packageFullPath := rootDir + "/" + packageAlias + "/"

	if mkErr := os.MkdirAll(packageFullPath, os.ModePerm); mkErr != nil {
		return mkErr
	}

	if addLibertyFile {
		var libertyData []byte
		if fileErr := ioutil.WriteFile(packageFullPath+utils.LibertyFile, libertyData, os.ModePerm); fileErr != nil {
			return fileErr
		}
	}

	return nil
}

func cleanupListTest() {
	os.RemoveAll(rootDir)
}
