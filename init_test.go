package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestCreatingLibertyFileWithValidData(t *testing.T) {
	defer cleanup()

	var data LibertyFile
	data.Project = Project{Author: "Some Guy", Description: "This is just a test"}

	GenerateLibertyFile(&data)

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

func cleanup() {
	os.Remove("liberty.json")
}
