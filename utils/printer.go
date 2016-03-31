package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

const (
	HeaderMessage = "===== Liberty ====="
)

// PrintHeader is used to print the name of the program in bold
func PrintHeader() {
	color.New(color.FgMagenta, color.Bold).Println(HeaderMessage)
}

// PrintInfo should be used to print any informative messages
func PrintInfo(info string) {
	color.Yellow(info)
}

// PrintUrgent should be used to print any errors or critical problems
func PrintUrgent(message string) {
	color.Red(message)
	PrintPadding()
}

// PrintSuccess should be used when an operation has been completed successfully
func PrintSuccess(message string) {
	color.Green(message)
}

// PrintClear clears the console screen
func PrintClear() {
	var cmd *exec.Cmd
	osType := runtime.GOOS

	if osType == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// PrintPadding will print 5 empty lines to the console
func PrintPadding() {
	fmt.Print("\n\n\n\n\n")
}
