package prompt

import (
	"fmt"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/pkg/constants"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
)

// Define constants for different stages of the prompt
const (
	stageLanguageSelection = iota
	stageProjectNameInput
	stageGoModulePathInput
	stageFrameworkSelection
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
	frameworkCursor   int
	frameworks        []contracts.FrameworkOption
	selectedFramework string
	inputError        string // inline validation error shown below the input
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
		if err := validateProjectName(m.projectName); err != nil {
			m.inputError = err.Error()
			return m, nil
		}
		m.inputError = ""

		// If Go is selected, move to module path input stage
		if m.choices[m.cursor] == constants.LanguageGo {
			m.frameworks = contracts.Frameworks[constants.LanguageGo]
			m.defaultModulePath = m.projectName
			m.modulePath = ""
			m.stage = stageGoModulePathInput
			return m, nil
		}
		if m.choices[m.cursor] == constants.LanguageNode {
			m.frameworks = contracts.Frameworks[constants.LanguageNode]
			m.stage = stageFrameworkSelection
			return m, nil
		}

		// For other languages, mark as done
		m.stage = stageFrameworkSelection
		m.done = true
		return m, tea.Quit

	case stageGoModulePathInput:
		// Use default if empty, otherwise validate what was typed
		if m.modulePath == "" {
			m.modulePath = m.defaultModulePath
		} else {
			if err := validateModulePath(m.modulePath); err != nil {
				m.inputError = err.Error()
				return m, nil
			}
		}
		m.inputError = ""
		m.frameworks = contracts.Frameworks[constants.LanguageGo]
		m.stage = stageFrameworkSelection
		return m, nil

	case stageFrameworkSelection:
		m.selectedFramework = m.frameworks[m.frameworkCursor].Value
		m.done = true
		return m, tea.Quit

	default:
		return m, nil
	}
}

// HandleMoveUpCase processes the cursor up action.
func HandleMoveUpCase(m BubbleTeaModel) BubbleTeaModel {
	if m.stage == stageFrameworkSelection {
		if m.frameworkCursor > 0 {
			m.frameworkCursor--
		}
		return m
	}

	if m.cursor > 0 {
		m.cursor--
	}
	return m
}

// HandleMoveDownCase processes the cursor down action.
func HandleMoveDownCase(m BubbleTeaModel) BubbleTeaModel {
	if m.stage == stageFrameworkSelection {
		if m.frameworkCursor < len(m.frameworks)-1 {
			m.frameworkCursor++
		}
		return m
	}

	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
	return m
}

// HandleBackspaceCase processes the backspace action.
func HandleBackspaceCase(m BubbleTeaModel) BubbleTeaModel {
	if m.stage == stageProjectNameInput && len(m.projectName) > 0 {
		m.projectName = m.projectName[:len(m.projectName)-1]
		m.inputError = ""
	}

	if m.stage == stageGoModulePathInput && len(m.modulePath) > 0 {
		m.modulePath = m.modulePath[:len(m.modulePath)-1]
		m.inputError = ""
	}
	return m
}

// HandleDefaultCase processes default character input.
// Only characters valid for the current field are accepted; invalid ones are silently dropped.
func HandleDefaultCase(m BubbleTeaModel, msg tea.KeyMsg) BubbleTeaModel {
	if len(msg.Runes) != 1 || !unicode.IsPrint(msg.Runes[0]) {
		return m
	}

	ch := msg.Runes[0]

	if m.stage == stageProjectNameInput {
		if isValidProjectNameChar(ch, len(m.projectName)) {
			m.projectName += string(ch)
			m.inputError = ""
		}
	}

	if m.stage == stageGoModulePathInput {
		if isValidModulePathChar(ch) {
			m.modulePath += string(ch)
			m.inputError = ""
		}
	}

	return m
}

// isValidProjectNameChar returns true if ch is allowed at the given position in a project name.
func isValidProjectNameChar(ch rune, pos int) bool {
	if unicode.IsLower(ch) || unicode.IsDigit(ch) {
		return true
	}
	if pos > 0 && (ch == '-' || ch == '_') {
		return true
	}
	return false
}

// isValidModulePathChar returns true if ch is allowed anywhere in a Go module path.
func isValidModulePathChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) ||
		ch == '-' || ch == '_' || ch == '.' || ch == '/'
}

// validateProjectName checks the full project name string against naming rules.
func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf(messages.EmptyProjectName)
	}
	runes := []rune(name)
	if !unicode.IsLower(runes[0]) {
		return fmt.Errorf(messages.ProjectNameLowerCase)
	}
	for _, ch := range runes[1:] {
		if !unicode.IsLower(ch) && !unicode.IsDigit(ch) && ch != '-' && ch != '_' {
			return fmt.Errorf(messages.ProjectNameHint)
		}
	}
	return nil
}

// validateModulePath checks the full Go module path string.
func validateModulePath(path string) error {
	if path == "" {
		return fmt.Errorf(messages.ModulePathEmpty)
	}
	if path[0] == '/' || path[len(path)-1] == '/' {
		return fmt.Errorf(messages.ModulePathHint)
	}
	for _, ch := range path {
		if !isValidModulePathChar(ch) {
			return fmt.Errorf(messages.ModulePathHint)
		}
	}
	return nil
}
