package main

import (
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
	"github.com/libgit2/git2go"
)

func TestSetGetsCorrectTag(t *testing.T) {
	repoPath := utils.CloneGitRepo("git://github.com/liberty-org/cli-test.git", "liberty-org/cli-test")
	args := []string{"liberty", "set", "liberty-org/cli-test", "v0.0.1"}

	ExecuteSet(args)

	repo, _ := git.OpenRepository(repoPath)
	isDetached, _ := repo.IsHeadDetached()

	if !isDetached {
		t.FailNow()
	}

	os.RemoveAll("./_libs")
}
