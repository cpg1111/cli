package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// PerformInit asks the user a number of questions to setup the project
// It creates a liberty (JSON) file with these settings (if provided)
func PerformInit() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("-- Welcome to Liberty! --")
	fmt.Println("We are going to ask you a few questions about your project.")
	fmt.Print("Is that okay? (Y/N): ")

	shouldAskQuestions := requestPermission(scanner)

	if !shouldAskQuestions {
		fmt.Print("You've now been liberated. Enjoy your freedom!")
		return
	}

	// Assuming yes here
	fmt.Print("Project Title: ")
	projectTitle := fetchInput(scanner)

	fmt.Print("Project Description: ")
	projectDesc := fetchInput(scanner)

	fmt.Print("Author Name: ")
	authorName := fetchInput(scanner)

	fmt.Printf("%s + %s + %s", projectTitle, projectDesc, authorName)
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
// It supports fuzzy responses (y/n, yes/no, YES/NO)
// It returns a boolean that represents the users response
func requestPermission(scanner *bufio.Scanner) bool {

	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}

	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	text := scanner.Text()
	if containsString(okayResponses, text) {
		return true
	} else if containsString(nokayResponses, text) {
		return false
	} else {
		fmt.Print("Please enter either [Y/N] to continue:")
		return requestPermission(scanner)
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
