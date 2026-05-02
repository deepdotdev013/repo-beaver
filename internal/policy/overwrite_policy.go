package policy

import (
	"fmt"
	"os"
	"strings"

	"github.com/deepdotdev013/repo-beaver/pkg/messages"
)

func AvoidDirOverwrite(path string) (bool, error) {
	// Check if the path exists
	info, err := os.Stat(path)

	// If it exists and is a directory, do not overwrite
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}

	// Path exists
	if !info.IsDir() {
		return false, fmt.Errorf(messages.PathExistsNotDirectory, path)
	}

	fmt.Printf(messages.OverwritePrompt, path)

	// Read user input
	var response string
	_, err = fmt.Scanln(&response)
	if err != nil {
		return false, err
	}

	// If the user does not confirm, do not overwrite
	if strings.ToLower(response) != "y" {
		return false, nil
	}

	// User confirmed, remove the existing directory
	if err := os.RemoveAll(path); err != nil {
		return false, err
	}

	return true, nil
}
