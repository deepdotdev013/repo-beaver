package generator

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/pkg/constants"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/templates"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

// Generator implements the Generator interface for Go projects.
type GoGenerator struct{}

// Generate creates the directory structure and main.go file for a Go project.
func (g *GoGenerator) Generate(cfg contracts.InitConfig) error {

	// Create directories
	for _, dir := range g.directories(cfg.ProjectName) {
		err := os.MkdirAll(filepath.Join(cfg.ProjectName, dir), 0755)
		if err != nil {
			return err
		}
		ui.LogStep(constants.LogStepCreated, filepath.Join(cfg.ProjectName, dir))
	}

	// Define template and destination paths
	files := []contracts.FileTemplate{
		{
			Tmpl: fmt.Sprintf("go/%s/main.go.tmpl", cfg.Framework),
			Dest: filepath.Join("cmd", cfg.ProjectName, "main.go"),
		},
		{
			Tmpl: fmt.Sprintf("go/%s/README.md.tmpl", cfg.Framework),
			Dest: "README.md",
		},
		{
			Tmpl: fmt.Sprintf("go/%s/gitignore.tmpl", cfg.Framework),
			Dest: ".gitignore",
		},
		{
			Tmpl: fmt.Sprintf("go/%s/workflow.yml.tmpl", cfg.Framework),
			Dest: ".github/workflows/ci.yml",
		},
		{
			Tmpl: fmt.Sprintf("go/%s/env.example.tmpl", cfg.Framework),
			Dest: ".env.example",
		},
		{
			Tmpl: fmt.Sprintf("go/%s/env.tmpl", cfg.Framework),
			Dest: ".env",
		},
		{
			Tmpl: fmt.Sprintf("go/%s/Dockerfile.tmpl", cfg.Framework),
			Dest: "Dockerfile",
		},
	}

	// Render the app.js template
	return templates.RenderFiles(cfg.ProjectName, files, templates.TemplateData{
		ProjectName: cfg.ProjectName,
	})

}

// Init initializes the Go project with necessary configurations.
func (g *GoGenerator) Init(cfg contracts.InitConfig) error {

	if cfg.ModulePath == "" {
		cfg.ModulePath = cfg.ProjectName
	}

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", cfg.ModulePath)

	// Set the command's working directory to the project directory
	cmd.Dir = cfg.ProjectName
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Run the command and handle any error encountered
	if err := cmd.Run(); err != nil {
		return err
	}
	ui.LogStep(constants.LogStepInit, "go.mod")

	// install framework
	switch cfg.Framework {

	case constants.FrameworkGin:
		cmd = exec.Command("go", "get", "github.com/gin-gonic/gin")

	case constants.FrameworkGorilla:
		cmd = exec.Command("go", "get", "github.com/gorilla/mux")

	case constants.FrameworkNone:
		return nil

	default:
		return fmt.Errorf("unsupported framework: %s", cfg.Framework)
	}

	if err := ui.RunSpinner(ui.Primary(fmt.Sprintf(messages.InstallingDependencies, cfg.Framework)), func() error {
		cmd.Dir = cfg.ProjectName
		cmd.Stdout = nil
		cmd.Stderr = io.Discard
		return cmd.Run()
	}); err != nil {
		return err
	}

	return nil
}

// directories returns the list of directories to be created for the Go project.
func (g *GoGenerator) directories(projectName string) []string {
	return []string{
		filepath.Join("cmd", projectName),
		"internal/handlers",
		"internal/services",
		"internal/repositories",
		"internal/models",
		"internal/domains",
		"internal/core",
		"pkg/logger",
		"pkg/utils",
		"configs",
		"tests",
		".github/workflows",
	}
}
