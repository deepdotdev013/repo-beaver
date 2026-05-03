package policy

import (
	"fmt"
	"os/exec"

	"github.com/deepdotdev013/repo-beaver/pkg/constants"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

// RequireCommand checks if a command is available in the system's PATH.
func RequireCommand(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%s not found in PATH", name)
	}
	return nil
}

// CheckLanguageDeps checks if the required dependencies for a given programming language are installed.
func CheckLanguageDeps(language string) error {
	switch language {

	case constants.LanguageGo:
		if err := RequireCommand(constants.LanguageGo); err != nil {
			return fmt.Errorf(
				ui.Error(messages.GoNotInstalled),
			)
		}

	case constants.LanguageNode:
		if err := RequireCommand(constants.LanguageNode); err != nil {
			return fmt.Errorf(
				ui.Error(messages.NodeNotInstalled),
			)
		}

		if err := RequireCommand("npm"); err != nil {
			return fmt.Errorf(
				ui.Error(messages.NpmNotInstalled),
			)
		}
	}

	return nil
}
