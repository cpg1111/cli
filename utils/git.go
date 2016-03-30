package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/libgit2/git2go"
)

type Repository struct {
	UrlPrefix    string
	ProviderFlag string
}

const (
	LibsDir      = "./_libs/"
	tagsPrefix   = "tags/"
	dotGitSuffix = ".git"
)

func TrimProviderPrefix(packageName string) string {
	repoDefs := ReadRepoDefinitions()

	var trimmedPackage string
	for _, repo := range repoDefs {
		if strings.HasPrefix(packageName, repo.ProviderFlag) {
			trimmedPackage = strings.TrimPrefix(packageName, repo.ProviderFlag)
		}
	}

	return trimmedPackage
}

// MakeGitPath creates a path to GitHub with a given string
// It will convert github: into git://github.com/
func MakeGitPath(rawPath string) (string, string) {
	var pathBuffer bytes.Buffer
	var userProjDir string

	repoDefs := ReadRepoDefinitions()

	for _, repo := range repoDefs {
		if strings.HasPrefix(rawPath, repo.ProviderFlag) {
			userProjDir = strings.TrimPrefix(rawPath, repo.ProviderFlag)

			pathBuffer.WriteString(repo.UrlPrefix)
			pathBuffer.WriteString(userProjDir)
			pathBuffer.WriteString(dotGitSuffix)
		}
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

func MakePathFromPackage(packageName string) string {
	path := LibsDir + packageName
	return path
}

func SetRepoVersion(packagePath string, packageVersion string) {
	repo, err := git.OpenRepository(packagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ref, err := repo.References.Dwim(packageVersion)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = repo.SetHead(ref.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
