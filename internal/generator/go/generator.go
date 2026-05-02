package generator

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/pkg/templates"
)

// Generator implements the Generator interface for Go projects.
type GoGenerator struct{}

// Generate creates the directory structure and main.go file for a Go project.
func (g *GoGenerator) Generate(projectName string) error {

	// Create directories
	for _, dir := range g.directories(projectName) {
		err := os.MkdirAll(filepath.Join(projectName, dir), 0755)
		if err != nil {
			return err
		}
	}

	// Define template and destination paths
	files := []contracts.FileTemplate{
		{
			Tmpl: "go/main.go.tmpl",
			Dest: filepath.Join("cmd", projectName, "main.go"),
		},
		{
			Tmpl: "go/README.md.tmpl",
			Dest: "README.md",
		},
		{
			Tmpl: "go/gitignore.tmpl",
			Dest: ".gitignore",
		},
	}

	// Render the app.js template
	return templates.RenderFiles(projectName, files, templates.TemplateData{
		ProjectName: projectName,
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

	// Run the command and return any error encountered
	return cmd.Run()
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
	}
}
