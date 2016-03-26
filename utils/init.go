package utils

import (
	"fmt"
	"log"
)

/* Perform the initialization process */
func Perform(args *string) {

	fmt.Println("-- Welcome to Liberty! --")
	fmt.Print("We are going to ask you a few questions about your project. Is that okay? (Y/N) ")

	// Ask for confirmation
	//response = fetchInput()

	// Assuming yes here
	fmt.Print("Project Title: ")
	var projectTitle = fetchInput()

	fmt.Print("Project Description: ")
	var projectDesc = fetchInput()

	fmt.Print("Author Name: ")
	var authorName = fetchInput()

	fmt.Printf("%s + %s + %s", projectTitle, projectDesc, authorName)
}

/** Fetches input and handles any errors */
func fetchInput() string {
	var response string

	_, err := fmt.Scanln(&response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}
