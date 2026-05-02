package ui

import (
	"github.com/fatih/color"
)

var (
	Primary = color.New(color.FgCyan, color.Bold).SprintFunc()
	Success = color.New(color.FgGreen).SprintFunc()
	Error   = color.New(color.FgRed).SprintFunc()
	Warning = color.New(color.FgYellow).SprintFunc()
	Muted   = color.New(color.FgHiBlack).SprintFunc()
	Bold    = color.New(color.Bold, color.FgGreen).SprintFunc()
	White   = color.New(color.FgWhite).SprintFunc()
)
