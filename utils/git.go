package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/libgit2/git2go"
)

type Repo struct {
	providerFlag string
	urlPrefix    string
}

type RepositoryList []Repo

const (
	LibsDir         = "./_libs/"
	githubFlag      = "github:"
	githubURLPrefix = "git://github.com/"
	dotGitSuffix    = ".git"
)

// MakeGitPath creates a path to GitHub with a given string
// It will convert github: into git://github.com/
func MakeGitPath(rawPath string) (string, string) {
	var pathBuffer bytes.Buffer
	var userProjDir string

	if strings.HasPrefix(rawPath, githubFlag) {
		userProjDir = strings.TrimPrefix(rawPath, githubFlag)

		pathBuffer.WriteString(githubURLPrefix)
		pathBuffer.WriteString(userProjDir)
		pathBuffer.WriteString(dotGitSuffix)
	}

	return pathBuffer.String(), userProjDir
}

// CloneGitRepo will clone a git repository at a given path.
// It will clone the project to the given destination
func CloneGitRepo(gitPath string, dest string) string {
	cloneOpts := &git.CloneOptions{}

	destPath := LibsDir + dest

	repo := make(chan string)

	go func() {
		repoOutput, err := git.Clone(gitPath, destPath, cloneOpts)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		repo <- repoOutput.Path()
	}()

	repository := <-repo

	return repository
}
