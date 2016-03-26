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

	if strings.HasPrefix(rawPath, githubFlag) {
		trimmedPath := strings.TrimPrefix(rawPath, githubFlag)

		pathBuffer.WriteString(githubUrlPrefix)
		pathBuffer.WriteString(trimmedPath)
		pathBuffer.WriteString(dotGitSuffix)
	}

	return pathBuffer.String()
}
