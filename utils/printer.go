package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/fatih/color"
)

const (
	HeaderMessage = "===== Liberty ====="
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// PrintHeader is used to print the name of the program in bold
func PrintHeader() {
	color.New(color.FgMagenta, color.Bold).Println(HeaderMessage)
}

// FPrintInfo should be used to print any formatted informative messages
func PrintInfof(message string, args ...interface{}) {
	color.Yellow(message, args)
}

// PrintInfo should be used to print any  informative messages
func PrintInfo(message string) {
	color.Yellow(message)
}

// PrintUrgent should be used to print any errors or critical problems
func PrintError(error string) {
	color.Red(error)
}

// PrintUrgentAndExit should be used to print any errors or critical problems.
// It also applys padding to the ouput and exits the program.
func PrintErrorThenExit(message string, errorCode int) {
	color.Red(message)
	PrintPadding()
	os.Exit(errorCode)
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

// GetConsoleWidth returns the number of characters that can fit across one line
func GetConsoleWidth() uint {
	width := getConsoleSize().Col
	return uint(width)
}

// GetConsoleHeight returns the number of lines that can fit in the console
func GetConsoleHeight() uint {
	height := getConsoleSize().Row
	return uint(height)
}

func getConsoleSize() winsize {
	ws := &winsize{}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		PrintErrorThenExit("Could not determine Console Size.", 1)
	}
	return *ws
}
