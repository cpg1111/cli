package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/libgit2/git2go"
)

const (
	libsDir         = "./_libs/"
	githubFlag      = "github:"
	githubUrlPrefix = "git://github.com/"
	dotGitSuffix    = ".git"
)

func MakeGitPath(rawPath string) (string, string) {
	var pathBuffer bytes.Buffer
	var userProjDir string

	if strings.HasPrefix(rawPath, githubFlag) {
		userProjDir = strings.TrimPrefix(rawPath, githubFlag)

		pathBuffer.WriteString(githubUrlPrefix)
		pathBuffer.WriteString(userProjDir)
		pathBuffer.WriteString(dotGitSuffix)
	}

	return pathBuffer.String(), userProjDir
}

func CloneGitRepo(gitPath string, dest string) string {
	cloneOpts := &git.CloneOptions{}

	destPath := libsDir + dest

	repo := make(chan string)

	go func() {
		repoOutput, err := git.Clone(gitPath, destPath, cloneOpts)
		repo <- repoOutput
	}()

	repository := <-repo

	fmt.Println(repository)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return repo.Path()
}
