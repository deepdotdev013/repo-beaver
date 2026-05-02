package generator

import (
	"fmt"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	goGen "github.com/deepdotdev013/repo-beaver/internal/generator/go"
	nodeGen "github.com/deepdotdev013/repo-beaver/internal/generator/node"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

type Generator interface {
	// Generate creates the project structure and necessary files.
	Generate(projectName string) error
	// Init initializes the project with necessary configurations.
	Init(cfg contracts.InitConfig) error
}

func Get(language string) (Generator, error) {
	switch language {
	case "go":
		return &goGen.GoGenerator{}, nil
	case "node":
		return &nodeGen.NodeGenerator{}, nil
	default:
		return nil, fmt.Errorf(ui.Error(messages.UnsupportedLanguageError), language)
	}
}
