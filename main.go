package main

import (
	"os"

	"github.com/liberty-org/cli/utils"
)

const (
	rootDir = "_libs"
)

func main() {
	utils.PrintClear()
	utils.PrintHeader()

	if !SubCommandsExist() {
		utils.PrintUrgent("Insufficient number of subcommands supplied")
		os.Exit(2)
	}

	ParseSubCommands(os.Args)
	utils.PrintPadding()
}
