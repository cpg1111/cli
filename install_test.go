package main

import (
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
)

func TestMakeGitPathForGithub(t *testing.T) {
	rawPath := "github:testuser/testrepo"

	gitPath, userProjDir := utils.MakeGitPath(rawPath)

	if gitPath != "git://github.com/testuser/testrepo.git" {
		t.FailNow()
	}

	if userProjDir != "testuser/testrepo" {
		t.FailNow()
	}
}

func TestGitClone(t *testing.T) {
	// Should find a better way to test this than running an actual clone
	repoPath := utils.CloneGitRepo("git://github.com/liberty-org/cli.git", "liberty-org/cli")

	if _, err := os.Stat(repoPath); err != nil {
		t.FailNow()
	}

	os.RemoveAll("./_libs")
}
