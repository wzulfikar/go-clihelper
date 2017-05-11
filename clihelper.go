package clihelper

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
)

var ExitHeader = "Whoops! Something went wrong :("
var ExitMsg = "\nPress enter to exit.."

func Exit(exitCode int) {
	Pause(ExitMsg)
	os.Exit(exitCode)
}

func ExitWithMessage(msg string, errorCode int) {
	if errorCode > 0 {
		oops := fmt.Sprintf("%s\n", color.RedString(ExitHeader))

		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, oops)
		} else {
			fmt.Printf(oops)
		}
	}

	fmt.Println(msg)
	Pause(ExitMsg)

	os.Exit(errorCode)
}

func CheckError(err error) {
	if err != nil {
		ExitWithMessage(err.Error(), 1)
	}
}

func Pause(msg string) {
	fmt.Println(msg)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
