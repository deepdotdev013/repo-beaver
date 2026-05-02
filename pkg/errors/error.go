package errors

import (
	"fmt"
	"os"
)

// HandleError prints the error message to stderr and exits the program with the specified exit code.
func HandleError(err error, code int) {
	// If there's no error, simply return.
	if err == nil {
		return
	}

	// Default exit code to 1 if not specified.
	exitCode := 1
	if exitCode != 1 {
		exitCode = code
	}

	// Print the error message to stderr and exit.
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(exitCode)
}
