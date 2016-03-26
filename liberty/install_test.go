package main

import (
	"os"
	"testing"
)

func TestMakeGitPathForGithub(t *testing.T) {
	rawPath := "github:testuser/testrepo"

	gitPath, userProjDir := MakeGitPath(rawPath)

	if gitPath != "git://github.com/testuser/testrepo.git" {
		t.FailNow()
	}

	if userProjDir != "testuser/testrepo" {
		t.FailNow()
	}
}

func TestGitClone(t *testing.T) {
	repoPath := CloneGitRepo("git://github.com/liberty-org/cli.git", "liberty-org/cli")

	if _, err := os.Stat(repoPath); err != nil {
		t.FailNow()
	}
}
