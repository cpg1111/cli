package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/liberty-org/cli/utils"
)

func TestCreatingLibertyFileWithValidData(t *testing.T) {
	defer cleanupInitTest()

	var data utils.LibertyData
	data.Title = "test"

	data.GenerateLibertyFile()

	if _, err := os.Stat("liberty.json"); os.IsNotExist(err) {
		t.FailNow()
	}
}

func TestUserPromptForYes(t *testing.T) {
	var input bytes.Buffer
	scanner := bufio.NewScanner(&input)

	input.WriteString("y")

	if !requestPermission(scanner) {
		t.FailNow()
	}
}

func TestUserPromptForNo(t *testing.T) {
	var input bytes.Buffer
	scanner := bufio.NewScanner(&input)

	input.WriteString("n")

	if requestPermission(scanner) {
		t.FailNow()
	}
}

func cleanupInitTest() {
	os.Remove("liberty.json")
}
