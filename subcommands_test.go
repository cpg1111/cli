package main

import (
	"os"
	"testing"
)

func TestGetSubCommand(t *testing.T) {
	args := []string{"liberty", "get", "github:liberty-org/cli"}
	localRepo := getCommand(args)

	if _, err := os.Stat(localRepo); err != nil {
		t.FailNow()
	}

	os.RemoveAll("./_libs")
}
