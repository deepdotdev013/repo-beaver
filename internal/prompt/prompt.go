package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

// initialModel initializes the Language prompt model.
func initialModel() BubbleTeaModel {
	return BubbleTeaModel{
		choices: []string{"go", "node"},
		stage:   stageLanguageSelection,
	}
}

// Init is the initial command for the model prompt.
func (m BubbleTeaModel) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages and updates the model's state accordingly.
func (m BubbleTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle key messages
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		// Handle case to quit the program
		case tea.KeyCtrlC, tea.KeyEsc:
			return HandleQuitCase(m)

		// Handle case to select a choice
		case tea.KeyEnter:
			return HandleSelectCase(m)

		// Handle cursor up movement
		case tea.KeyUp:
			m = HandleMoveUpCase(m)
		// Handle cursor down movement
		case tea.KeyDown:
			m = HandleMoveDownCase(m)

		case tea.KeyBackspace:
			m = HandleBackspaceCase(m)

		// Handle character input for project name
		default:
			m = HandleDefaultCase(m, msg)
		}

	}
	return m, nil
}

// View renders the current state of the model as a string.
func (m BubbleTeaModel) View() string {

	// If the model state is done, return an empty string
	if m.done {
		return ""
	}

	// Language selection stage
	if m.stage == stageLanguageSelection {
		s := ui.Primary(messages.SelectBackendLanguage)

		// Render the list of choices
		for i, choice := range m.choices {
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ui.Success("➜")
				choice = ui.Bold(choice)
			}
			// Add the choice to the string
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}

		s += ui.Muted(messages.LanguageNavigationHelp)

		return s + ui.Muted(messages.QuitHint)
	}

	if m.stage == stageProjectNameInput {
		// Return project name input stage
		return fmt.Sprintf(
			messages.EnterProjectNamePrompt,
			m.projectName,
		) + "\n"
	}

	if m.stage == stageGoModulePathInput {
		return fmt.Sprintf(
			messages.GoModulePathPrompt,
			m.modulePath,
			m.defaultModulePath,
		) + "\n"
	}

	return ""
}

// StartLanguagePrompt initiates the Bubble tea TUI and returns the selected project name and language.
func StartLanguagePrompt() (string, string, string, error) {
	// Create a new Bubble tea program with the initial model
	program := tea.NewProgram(initialModel())

	// Start the program and wait for it to finish
	m, err := program.Run()
	if err != nil {
		return "", "", "", err
	}

	// Type assert the returned model to our model type
	finalModel := m.(BubbleTeaModel)

	if finalModel.contextCancelled {
		return "", "", "", fmt.Errorf(messages.ErrPromptCancelled)
	}

	if finalModel.projectName == "" {
		return "", "", "", fmt.Errorf(messages.EmptyProjectName)
	}

	// Return the selected project name and language
	return finalModel.projectName, finalModel.choices[finalModel.cursor], finalModel.modulePath, nil
}
