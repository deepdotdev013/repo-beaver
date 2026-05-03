package messages

const (
	// Success and Warning Messages
	ProjectGeneratedSuccess  string = "\n✔ Done in %.2f seconds - Let’s build something meaningful. Happy hacking! :)\n"
	ErrPromptCancelled       string = "⚠️ Operation cancelled by user"
	CreatingProjectStructure string = "Hi! Scaffolding your backend...\n"
	InitializingProject      string = "Initializing project"
	InstallingDependencies   string = "Installing %s framework dependencies"

	// Validation Messages
	EmptyProjectName     string = "⚠️ Project name cannot be empty"
	ProjectNameHint      string = "Allowed: lowercase letters, digits, hyphens (-), underscores (_). Must start with a letter."
	ProjectNameLowerCase string = "Project name must start with a lowercase letter"
	ModulePathHint       string = "Allowed: letters, digits, hyphens (-), underscores (_), dots (.), slashes (/). e.g. github.com/username/project"
	ModulePathEmpty      string = "⚠️ Module path cannot be empty"

	// Prompt Messages
	SelectBackendLanguage  string = "Select the backend language for your project:\n\n"
	LanguageNavigationHelp string = "\nUse ↑ / ↓ arrow keys to navigate and press Enter to confirm."
	QuitHint               string = "\nPress Esc or Ctrl+C at any time to cancel the operation and exit.\n"
	GoModulePathPrompt            = "Go requires a module path to manage imports and dependencies.\n" +
		"It is recommended to use a full path (e.g., github.com/username/project-name) " +
		"to avoid refactoring later when the project is shared or deployed.\n\n" +
		"Enter Go module path:"
	EnterProjectNamePrompt string = "Enter a name for your project (this will be used as the folder name):\n\n"
	PressEnterToContinue   string = "Press Enter to continue."
	PressEnterToUseDefault string = "Press Enter to use the default value."

	// Informational Messages
	GoNotInstalled string = "Go is not installed or not available in PATH.\n\n" +
		"This project requires Go to generate and initialize files.\n" +
		"Download and install it from: https://go.dev/dl/\nThen re-run this command."

	NodeNotInstalled string = "Node.js is not installed or not available in PATH.\n\n" +
		"This project requires:\n" +
		"  - node (>= 18)\n" +
		"  - npm\n\n" +
		"Node.js is required to run the application and manage dependencies.\n\n" +
		"Install it from: https://nodejs.org/en/download\n\n" +
		"Or install it using a version manager: https://github.com/nvm-sh/nvm\n\n" +
		"After installation, re-run this command."

	NpmNotInstalled string = "npm is not installed or not available in PATH.\n\n" +
		"npm usually comes bundled with Node.js.\n" +
		"Try reinstalling Node.js from: https://nodejs.org/\nThen re-run this command."

	SelectFrameworkPrompt string = "Select a framework for your project:\n\n"

	// Filesystem Messages
	PathExistsNotDirectory string = "Cannot continue: \"%s\" exists but is not a directory."
	OverwritePrompt        string = "Directory \"%s\" already exists.\nOverwriting will delete all existing files.\nDo you want to continue? (y/N): "
	OverwriteCancelled     string = "Operation cancelled. Your existing directory was left unchanged."

	// Language & Dependency Errors
	UnsupportedLanguageError string = "Unsupported language selected: %s"
)
