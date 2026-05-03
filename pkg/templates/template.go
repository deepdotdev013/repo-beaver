package templates

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/pkg/constants"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

//go:embed node/* go/*
var FS embed.FS

type TemplateData struct {
	ProjectName string
}

// RenderTemplate renders a template and writes it safely
func RenderTemplate(tmplPath, destPath string, data TemplateData) error {
	// Parse template from embedded FS
	t, err := template.ParseFS(FS, tmplPath)
	if err != nil {
		return err
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// Create file
	f, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Execute template
	return t.Execute(f, data)
}

// RenderFiles renders multiple templates to their respective destination paths.
func RenderFiles(projectName string, files []contracts.FileTemplate, data TemplateData) error {
	for _, file := range files {
		fullPath := filepath.Join(projectName, file.Dest)
		err := RenderTemplate(
			file.Tmpl,
			fullPath,
			data,
		)
		if err != nil {
			return err
		}

		ui.LogStep(constants.LogStepWrite, fullPath)
	}
	return nil
}
