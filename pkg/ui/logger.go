package ui

import (
	"fmt"
	"strings"
)

// logBuffer is set by RunSpinner to capture LogStep output during spinner execution.
// When nil, LogStep writes directly to stdout.
var logBuffer *[]string

// Logger functions for consistent output formatting across the application.
func LogStep(action, path string) {
	const width = 12
	padding := width - len([]rune(action))
	if padding < 0 {
		padding = 0
	}
	line := fmt.Sprintf(Primary("✓ %s%s %s\n"), action, strings.Repeat(" ", padding), path)

	if logBuffer != nil {
		*logBuffer = append(*logBuffer, line)
	} else {
		fmt.Print(line)
	}
}
