package main

import (
	"bytes"
	"fmt"
	"github.com/libgit2/git2go"
	"os"
	"strings"
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
		repoOutput, _ := git.Clone(gitPath, destPath, cloneOpts)
		repo <- repoOutput
	}()

	repositoty := <-repo

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return repository.Path()
}
