package main

import (
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
)

func TestRemoveDeletesPackage(t *testing.T) {
	os.MkdirAll("./_libs/testA/testB", os.ModePerm)

	args := []string{"liberty", "remove", "github:testA/testB"}

	libertyData := &utils.LibertyData{Title: "TestRemoveDeletesPackage"}
	testDep := &utils.Dependency{Name: "github:dont/delete"}
	dep := &utils.Dependency{Name: "github:testA/testB"}

	libertyData.Dependencies = append(libertyData.Dependencies, *dep)
	libertyData.Dependencies = append(libertyData.Dependencies, *testDep)
	libertyData.GenerateLibertyFile()

	ExecuteRemove(args)

	if _, err := os.Stat("./_libs/testA"); err == nil {
		t.FailNow()
	}

	os.Remove("./liberty.json")
	os.RemoveAll("./_libs")
}
