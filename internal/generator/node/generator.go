package generator

import (
	"encoding/json"
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

// Generator implements the Generator interface for Node.js projects.
type NodeGenerator struct{}

// Generate creates the directory structure and app.js file for a Node.js project.
func (n *NodeGenerator) Generate(cfg contracts.InitConfig) error {

	// Create directories
	for _, dir := range n.directories() {
		err := os.MkdirAll(filepath.Join(cfg.ProjectName, dir), 0755)
		if err != nil {
			return err
		}
		ui.LogStep(constants.LogStepCreated, filepath.Join(cfg.ProjectName, dir))
	}

	files := []contracts.FileTemplate{
		{
			Tmpl: fmt.Sprintf("node/%s/app.js.tmpl", cfg.Framework),
			Dest: "app.js",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/README.md.tmpl", cfg.Framework),
			Dest: "README.md",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/gitignore.tmpl", cfg.Framework),
			Dest: ".gitignore",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/workflow.yml.tmpl", cfg.Framework),
			Dest: ".github/workflows/ci.yml",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/env.example.tmpl", cfg.Framework),
			Dest: ".env.example",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/env.tmpl", cfg.Framework),
			Dest: ".env",
		},
		{
			Tmpl: fmt.Sprintf("node/%s/Dockerfile.tmpl", cfg.Framework),
			Dest: "Dockerfile",
		},
	}

	// Render the app.js template
	return templates.RenderFiles(cfg.ProjectName, files, templates.TemplateData{
		ProjectName: cfg.ProjectName,
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

	// Run the command and handle any error encountered
	if err := cmd.Run(); err != nil {
		return err
	}
	ui.LogStep(constants.LogStepInit, "package.json")
	switch cfg.Framework {
	case constants.FrameworkExpress:
		cmd = exec.Command("npm", "install", "express")

	case constants.FrameworkFastify:
		cmd = exec.Command("npm", "install", "fastify")

	case constants.FrameworkNone:
		// no dependency

	default:
		return fmt.Errorf("unsupported framework: %s", cfg.Framework)
	}

	if cfg.Framework != constants.FrameworkNone {
		if err := ui.RunSpinner(ui.Primary(fmt.Sprintf(messages.InstallingDependencies, cfg.Framework)), func() error {
			cmd.Dir = cfg.ProjectName
			cmd.Stdout = nil
			cmd.Stderr = io.Discard
			return cmd.Run()
		}); err != nil {
			return err
		}
	}

	// Step 3: update package.json
	return updatePackageJSON(cfg)
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
		"src/validators",
		"configs",
		"tests",
		".github/workflows",
	}
}

// updatePackageJSON reads the existing package.json file, updates the scripts section, and writes it back to package.json.
func updatePackageJSON(cfg contracts.InitConfig) error {
	// Read the existing package.json file
	path := filepath.Join(cfg.ProjectName, "package.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a map, update the scripts section, and write it back to package.json
	var pkg map[string]interface{}
	json.Unmarshal(data, &pkg)

	pkg["scripts"] = map[string]string{
		"start": "node app.js",
	}

	newData, _ := json.MarshalIndent(pkg, "", "  ")
	return os.WriteFile(path, newData, 0644)
}
