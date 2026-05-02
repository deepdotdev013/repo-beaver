package generator

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/pkg/templates"
)

// Generator implements the Generator interface for Node.js projects.
type NodeGenerator struct{}

// Generate creates the directory structure and app.js file for a Node.js project.
func (n *NodeGenerator) Generate(projectName string) error {

	// Create directories
	for _, dir := range n.directories() {
		err := os.MkdirAll(filepath.Join(projectName, dir), 0755)
		if err != nil {
			return err
		}
	}

	files := []contracts.FileTemplate{
		{
			Tmpl: "node/app.js.tmpl",
			Dest: "app.js",
		},
		{
			Tmpl: "node/README.md.tmpl",
			Dest: "README.md",
		},
		{
			Tmpl: "node/gitignore.tmpl",
			Dest: ".gitignore",
		},
	}

	// Render the app.js template
	return templates.RenderFiles(projectName, files, templates.TemplateData{
		ProjectName: projectName,
	})
}

// Init initializes the Node.js project with necessary configurations.
func (g *NodeGenerator) Init(cfg contracts.InitConfig) error {
	// Create the project directory if it doesn't exist
	if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
		return err
	}

	// Initialize Node.js project with default settings
	cmd := exec.Command("npm", "init", "-y")

	// Set the command's working directory to the project directory
	cmd.Dir = cfg.ProjectName
	cmd.Stdout = nil
	cmd.Stderr = nil

	// Run the command and return any error encountered
	return cmd.Run()
}

// directories returns the list of directories to be created for the Node.js project.
func (n *NodeGenerator) directories() []string {
	return []string{
		"src",
		"src/models",
		"src/controllers",
		"src/services",
		"src/repositories",
		"src/routes",
		"src/middlewares",
		"src/utils",
		"src/policies",
		"configs",
	}
}
