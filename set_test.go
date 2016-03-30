package main

import (
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
	"github.com/libgit2/git2go"
)

func TestSetGetsCorrectTag(t *testing.T) {
	libertyData := &utils.LibertyData{}

	dep := &utils.Dependency{Name: "github:liberty-org/cli-test", Version: ""}
	libertyData.Dependencies = append(libertyData.Dependencies, *dep)

	libertyData.GenerateLibertyFile()

	repoPath := utils.CloneGitRepo("git://github.com/liberty-org/cli-test.git", "liberty-org/cli-test")
	args := []string{"liberty", "set", "github:liberty-org/cli-test", "v0.0.1"}

	ExecuteSet(args)

	repo, _ := git.OpenRepository(repoPath)
	isDetached, _ := repo.IsHeadDetached()

	if !isDetached {
		t.FailNow()
	}

	libFile := utils.ReadLibertyData()

	if libFile.Dependencies[0].Version != "v0.0.1" {
		t.FailNow()
	}

	os.Remove("./liberty.json")
	os.RemoveAll("./_libs")
}
