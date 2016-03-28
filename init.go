package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/liberty-org/cli/utils"
)

// ExecuteInit asks the user a number of questions to setup the project
// It creates a liberty (JSON) file with these settings (if provided)
func ExecuteInit(args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	var libertyData utils.LibertyData

	fmt.Println("*** Welcome to Liberty! ***")
	fmt.Println("We are going to ask you a few questions about your project.")
	fmt.Print("Is that okay? (Y/N): ")

	shouldAskQuestions := requestPermission(scanner)

	if shouldAskQuestions {
		fmt.Print("Project Title: ")
		libertyData.Title = fetchInput(scanner)

		fmt.Print("Project Description: ")
		libertyData.Description = fetchInput(scanner)

		fmt.Print("Author Name: ")
		libertyData.Author = fetchInput(scanner)

		libertyData.Dependencies = make([]utils.Dependency, 1)
	}

	libertyData.GenerateLibertyFile()
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
