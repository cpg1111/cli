package main

import (
	"os"
	"testing"
)

func TestGetSubCommand(t *testing.T) {
	args := []string{"liberty", "get", "github:liberty-org/cli"}
	libertyData := &LibertyData{Title: "TestGetAddsPackage"}
	dep := &Dependency{Name: "github:testA/testB"}

	libertyData.Dependencies = append(libertyData.Dependencies, *dep)
	libertyData.GenerateLibertyFile()

	localRepo := getCommand(args)

	if _, err := os.Stat(localRepo); err != nil {
		t.FailNow()
	}

	os.Remove("./liberty.json")
	os.RemoveAll("./_libs")
}
