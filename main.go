package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cmd := os.Args[1]  // Program Name is always the first (implicit) argument

	if cmd == "help" {
		for _, command := range GetHelpArguments() {
			fmt.Println()
			fmt.Printf("---%s---", command.Name)
			fmt.Println()
			fmt.Println(strings.TrimSpace(command.Usage))
		}
	}

	fmt.Printf("Program Name: %s\n", cmd)
}
