package ui

import (
	"fmt"

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
func RunSpinner(message string, task func() error) error {
	model := NewSpinner(message)
	p := tea.NewProgram(model)

	errChan := make(chan error, 1) // 👈 buffered channel (important)

	go func() {
		err := task()
		errChan <- err
		p.Send(DoneMsg{})
	}()

	_, err := p.Run()
	if err != nil {
		return err
	}

	return <-errChan
}
