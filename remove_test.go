package main

import (
	"os"
	"testing"
)

func TestRemoveDeletesPackage(t *testing.T) {
	os.MkdirAll("./_libs/testA/testB", os.ModePerm)

	args := []string{"liberty", "remove", "github:testA/testB"}

	ExecuteRemove(args)

	if _, err := os.Stat("./_libs/testA"); err == nil {
		t.FailNow()
	}

	os.RemoveAll("./_libs")
}
