package policy

import (
	"fmt"
	"os/exec"

	"github.com/deepdotdev013/repo-beaver/pkg/messages"
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

	case "go":
		if err := RequireCommand("go"); err != nil {
			return fmt.Errorf(
				messages.GoNotInstalled,
			)
		}

	case "node":
		if err := RequireCommand("node"); err != nil {
			return fmt.Errorf(
				messages.NodeNotInstalled,
			)
		}

		if err := RequireCommand("npm"); err != nil {
			return fmt.Errorf(
				messages.NpmNotInstalled,
			)
		}
	}

	return nil
}
