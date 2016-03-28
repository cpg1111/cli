package main

import (
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
)

func TestGetSubCommand(t *testing.T) {
	args := []string{"liberty", "get", "github:liberty-org/cli"}
	libertyData := &utils.LibertyData{Title: "TestGetAddsPackage"}
	dep := &utils.Dependency{Name: "github:testA/testB"}

	libertyData.Dependencies = append(libertyData.Dependencies, *dep)
	libertyData.GenerateLibertyFile()

	localRepo := getCommand(args)

	if _, err := os.Stat(localRepo); err != nil {
		t.FailNow()
	}

	os.Remove("./liberty.json")
	os.RemoveAll("./_libs")
}
