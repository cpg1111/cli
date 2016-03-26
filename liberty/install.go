package main

import (
	"bytes"
	"strings"
)

const (
	githubFlag      = "github:"
	githubUrlPrefix = "https://github.com/"
	dotGitSuffix    = ".git"
)

func MakeGitPath(rawPath string) string {
	var pathBuffer bytes.Buffer

	trimmedPath := strings.TrimPrefix(rawPath, githubFlag)

	if strings.HasPrefix(rawPath, githubFlag) {
		pathBuffer.WriteString(githubUrlPrefix)
		pathBuffer.WriteString(trimmedPath)
		pathBuffer.WriteString(dotGitSuffix)
	}

	return pathBuffer.String()
}
