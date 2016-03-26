package main

import (
	"testing"
)

func TestMakeGitPathForGithub(t *testing.T) {
	rawPath := "github:testuser/testrepo"

	gitPath := MakeGitPath(rawPath)

	if gitPath != "https://github.com/testuser/testrepo.git" {
		t.FailNow()
	}
}
