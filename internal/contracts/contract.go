package contracts

import "github.com/deepdotdev013/repo-beaver/pkg/constants"

// InitConfig represents the configuration for initializing a new project.
type InitConfig struct {
	ProjectName string
	ModulePath  string
	Framework   string
}

// FileTemplate represents a template for a file that can be generated during project initialization.
type FileTemplate struct {
	Tmpl string
	Dest string
}

// FrameworkOption represents a framework option that can be selected during project initialization.
type FrameworkOption struct {
	Name        string
	Description string
	Value       string
}

// Frameworks is a map of programming languages to their respective framework options.
var Frameworks = map[string][]FrameworkOption{
	constants.LanguageNode: {
		{"Express", "Minimal, most popular", constants.FrameworkExpress},
		{"Fastify", "High performance, modern", constants.FrameworkFastify},
		{"None", "Bare Node.js project", constants.FrameworkNone},
	},
	constants.LanguageGo: {
		{"Gin", "Fast, popular web framework", constants.FrameworkGin},
		{"Gorilla Mux", "Classic routing library", constants.FrameworkGorilla},
		{"None", "Standard net/http", constants.FrameworkNone},
	},
}
