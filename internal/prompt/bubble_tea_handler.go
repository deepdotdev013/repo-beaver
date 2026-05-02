package prompt

import (
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

// Define constants for different stages of the prompt
const (
	stageLanguageSelection = iota
	stageProjectNameInput
	stageGoModulePathInput
)

// BubbleTeaModel represents the state of the language selection prompt.
type BubbleTeaModel struct {
	cursor            int
	choices           []string // go, node
	projectName       string
	modulePath        string
	defaultModulePath string
	stage             int // Stages from the constants above
	done              bool
	contextCancelled  bool // indicates if the user context was cancelled
}

// --- Handler Functions ---

// HandleQuitCase processes the quit action.
func HandleQuitCase(m BubbleTeaModel) (tea.Model, tea.Cmd) {
	m.contextCancelled = true
	return m, tea.Quit
}

// HandleSelectCase processes the selection action based on the current stage.
func HandleSelectCase(m BubbleTeaModel) (tea.Model, tea.Cmd) {
	// Handle selection based on the current stage
	switch m.stage {

	// Language selection stage
	case stageLanguageSelection:
		m.stage = stageProjectNameInput
		return m, nil

		// Project name input stage
	case stageProjectNameInput:
		// If Go is selected, move to module path input stage
		if m.choices[m.cursor] == "go" {
			// default module path
			m.defaultModulePath = m.projectName
			m.modulePath = ""
			m.stage = stageGoModulePathInput
			return m, nil
		}

		// For other languages, mark as done
		m.done = true
		return m, tea.Quit

	case stageGoModulePathInput:
		if m.modulePath == "" {
			m.modulePath = m.defaultModulePath
		}

		m.done = true
		return m, tea.Quit

	default:
		return m, nil
	}
}

// HandleMoveUpCase processes the cursor up action.
func HandleMoveUpCase(m BubbleTeaModel) BubbleTeaModel {
	if m.cursor > 0 {
		m.cursor--
	}
	return m
}

// HandleMoveDownCase processes the cursor down action.
func HandleMoveDownCase(m BubbleTeaModel) BubbleTeaModel {
	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
	return m
}

// HandleBackspaceCase processes the backspace action.
func HandleBackspaceCase(m BubbleTeaModel) BubbleTeaModel {
	if m.stage == stageProjectNameInput && len(m.projectName) > 0 {
		m.projectName = m.projectName[:len(m.projectName)-1]
	}

	if m.stage == stageGoModulePathInput && len(m.modulePath) > 0 {
		m.modulePath = m.modulePath[:len(m.modulePath)-1]
	}
	return m
}

// HandleDefaultCase processes default character input.
func HandleDefaultCase(m BubbleTeaModel, msg tea.KeyMsg) BubbleTeaModel {
	// Here, we only accept printable characters for the project name
	if len(msg.Runes) == 1 && unicode.IsPrint(msg.Runes[0]) {
		if m.stage == stageProjectNameInput {
			m.projectName += string(msg.Runes[0])
		}

		if m.stage == stageGoModulePathInput {
			m.modulePath += string(msg.Runes[0])
		}
	}
	return m
}
