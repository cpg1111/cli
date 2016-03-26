package main

import (
	"strings"
	"testing"
)

func TestMakeGitPathForGithub(t *testing.T) {
	rawPath := "github:testuser/testrepo"

	gitPath := MakeGitPath(rawPath)

	if !strings.HasSuffix(gitPath, "https://github.com/testuser/testrepo.git") {
		t.FailNow()
	}
}
