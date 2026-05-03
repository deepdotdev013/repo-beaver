package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type SpinnerModel struct {
	spinner spinner.Model
	message string
	done    bool
}

func NewSpinner(message string) SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return SpinnerModel{
		spinner: s,
		message: message,
	}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case DoneMsg:
		m.done = true
		return m, tea.Quit
	}

	return m, nil
}

func (m SpinnerModel) View() string {
	if m.done {
		return fmt.Sprintf("✔ %s\n", m.message)
	}
	return fmt.Sprintf("%s %s...", m.spinner.View(), m.message)
}

// DoneMsg is used to stop spinner
type DoneMsg struct{}

// RunSpinner runs a spinner with the given message while executing the provided task function.
// Any LogStep calls made inside task are buffered and printed after the spinner exits,
// so bubbletea's terminal control does not mangle them.
func RunSpinner(message string, task func() error) error {
	model := NewSpinner(message)
	p := tea.NewProgram(model)

	errChan := make(chan error, 1)
	var logLines []string

	// Redirect LogStep output into a buffer while the spinner is running.
	logBuffer = &logLines

	go func() {
		err := task()
		errChan <- err
		p.Send(DoneMsg{})
	}()

	_, err := p.Run()

	// Restore normal LogStep output before printing buffered lines.
	logBuffer = nil

	if err != nil {
		return err
	}

	taskErr := <-errChan

	// Print all buffered log lines now that the terminal is ours again.
	fmt.Print(strings.Join(logLines, ""))

	return taskErr
}
