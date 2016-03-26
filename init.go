package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	libertyFile = "./liberty.json"
)

// ExecuteInit asks the user a number of questions to setup the project
// It creates a liberty (JSON) file with these settings (if provided)
func ExecuteInit(args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	var libertyData LibertyFile
	var projectData Project

	fmt.Println("*** Welcome to Liberty! ***")
	fmt.Println("We are going to ask you a few questions about your project.")
	fmt.Print("Is that okay? (Y/N): ")

	shouldAskQuestions := requestPermission(scanner)

	if shouldAskQuestions {

		fmt.Print("Project Title: ")
		projectData.Title = fetchInput(scanner)

		fmt.Print("Project Description: ")
		projectData.Description = fetchInput(scanner)

		fmt.Print("Author Name: ")
		projectData.Author = fetchInput(scanner)

		libertyData.Project = projectData
	}

	GenerateLibertyFile(&libertyData)
}

// Generates the Liberty File with the given data
// It saves it as "liberty.json" in the directory that the program was executed
func GenerateLibertyFile(data *LibertyFile) {
	jsonData, marshalErr := json.MarshalIndent(data, "", "\t")
	if marshalErr != nil {
		panic(marshalErr)
	}

	if err := ioutil.WriteFile(libertyFile, jsonData, 0644); err != nil {
		panic(err)
	}
}

// Fetches input and handles any errors
func fetchInput(scanner *bufio.Scanner) string {
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}

// Ask the user for a YES/NO response
// If the users response starts with y, the return value is true, else, false.
// It returns a boolean that represents the users response
func requestPermission(scanner *bufio.Scanner) bool {
	text := fetchInput(scanner)

	if len(text) > 0 && strings.ToUpper(text[:1]) == "Y" {
		return true
	}

	return false
}
