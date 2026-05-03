package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/internal/generator"
	"github.com/deepdotdev013/repo-beaver/internal/policy"
	"github.com/deepdotdev013/repo-beaver/internal/prompt"
	"github.com/deepdotdev013/repo-beaver/pkg/constants"
	"github.com/deepdotdev013/repo-beaver/pkg/errors"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
	"github.com/spf13/cobra"
)

// Flags for the create command
var (
	projectName    string
	langFlag       string
	frameworkFlag  string
	expressFlag    bool
	fastifyFlag    bool
	ginFlag        bool
	gorillaMuxFlag bool
	cfg            contracts.InitConfig
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new backend project",
	Args:  cobra.MaximumNArgs(1),

	// Reject duplicate flags before any business logic runs.
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkDuplicateFlags(os.Args[1:])
	},

	// The actual implementation of project creation will go here
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.Primary("🦫 Repo Beaver"))
		fmt.Println(ui.White("Generate backend projects in seconds.\n"))

		// Step 1: Handle flags to determine language and framework if not already set by prompts.
		var flagErr error
		langFlag, frameworkFlag, flagErr = handleFlags(frameworkFlag, langFlag, expressFlag, fastifyFlag, ginFlag, gorillaMuxFlag)
		errors.HandleError(flagErr, 1)

		// If a project name is provided as an argument, use it. Otherwise, it will be prompted later.
		if len(args) > 0 {
			projectName = args[0]
		}

		// If the language flag is provided, we can skip the interactive prompts and directly use the flags to set up the configuration.
		if langFlag != "" {
			cfg = contracts.InitConfig{
				ProjectName: projectName,
				Language:    langFlag,
				Framework:   frameworkFlag,
			}
		} else {
			// If no language flag is provided, we need to prompt the user for all necessary information.
			// Pass the project name (may be empty) so the prompt can skip the name stage if already supplied.
			projectName, language, modulePath, framework, err := prompt.StartLanguagePrompt(projectName)
			errors.HandleError(err, 1)

			// Set up the configuration based on user input.
			cfg = contracts.InitConfig{
				ProjectName: projectName,
				Language:    language,
				ModulePath:  modulePath,
				Framework:   framework,
			}
		}

		// Step 2: Check for language-specific dependencies.
		policyErr := policy.CheckLanguageDeps(cfg.Language)
		if policyErr != nil {
			errors.HandleError(policyErr, 1)
		}

		// Step 3: Check for directory overwrite policy.
		ok, err := policy.AvoidDirOverwrite(cfg.ProjectName)
		errors.HandleError(err, 1)
		if !ok {
			fmt.Println(messages.OverwriteCancelled)
			return
		}

		// Step 4: Get the appropriate generator based on the selected language.
		gen, err := generator.Get(cfg.Language)
		errors.HandleError(err, 1)

		// Record the start time for performance measurement.
		start := time.Now()

		// Step 5: Generate the project structure.
		err = ui.RunSpinner(ui.Primary(messages.CreatingProjectStructure), func() error {
			return gen.Generate(cfg)
		})
		errors.HandleError(err, 1)

		// Step 6: Initialize the project with necessary configurations.
		err = ui.RunSpinner(ui.Primary(messages.InitializingProject), func() error {
			return gen.Init(cfg)
		})
		errors.HandleError(err, 1)

		duration := time.Since(start)
		fmt.Println(ui.Success(fmt.Sprintf(messages.ProjectGeneratedSuccess, duration.Seconds())))
		fmt.Println(ui.Success(messages.NextSteps(cfg.ProjectName, cfg.Language, cfg.Framework)))
	},
}

// Initialize the create command and add it to the root command
func init() {
	createCmd.Flags().StringVar(&langFlag, "lang", "", "Language (go/node)")
	createCmd.Flags().StringVar(&frameworkFlag, "framework", "", "Framework (gin/gorilla mux/express/fastify)")
	createCmd.Flags().BoolVar(&expressFlag, constants.FrameworkExpress, false, "Express.js framework")
	createCmd.Flags().BoolVar(&fastifyFlag, constants.FrameworkFastify, false, "Fastify framework")
	createCmd.Flags().BoolVar(&ginFlag, constants.FrameworkGin, false, "Gin framework")
	createCmd.Flags().BoolVar(&gorillaMuxFlag, constants.FrameworkGorilla, false, "Gorilla Mux framework")

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
}

// handleFlags processes the command-line flags to determine the language and framework for project generation.
// Returns an error if more than one framework shortcut flag is provided at the same time.
func handleFlags(frameworkFlag string, langFlag string, expressFlag bool, fastifyFlag bool, ginFlag bool, gorillaMuxFlag bool) (string, string, error) {
	// Count how many shortcut flags were set simultaneously.
	shortcutCount := 0
	if expressFlag {
		shortcutCount++
	}
	if fastifyFlag {
		shortcutCount++
	}
	if ginFlag {
		shortcutCount++
	}
	if gorillaMuxFlag {
		shortcutCount++
	}

	if shortcutCount > 1 {
		return "", "", fmt.Errorf(ui.Error(messages.ConflictingFrameworkFlags))
	}

	if frameworkFlag != "" && langFlag == "" {
		switch frameworkFlag {
		case constants.FrameworkExpress, constants.FrameworkFastify:
			langFlag = constants.LanguageNode
		case constants.FrameworkGin, constants.FrameworkGorilla:
			langFlag = constants.LanguageGo
		}
	} else if expressFlag {
		langFlag = constants.LanguageNode
		frameworkFlag = constants.FrameworkExpress
	} else if fastifyFlag {
		langFlag = constants.LanguageNode
		frameworkFlag = constants.FrameworkFastify
	} else if ginFlag {
		langFlag = constants.LanguageGo
		frameworkFlag = constants.FrameworkGin
	} else if gorillaMuxFlag {
		langFlag = constants.LanguageGo
		frameworkFlag = constants.FrameworkGorilla
	}
	return langFlag, frameworkFlag, nil
}

// checkDuplicateFlags scans raw CLI args for any flag that appears more than once.
func checkDuplicateFlags(args []string) error {
	seen := make(map[string]int)
	for _, arg := range args {
		// Normalise --flag and -flag to just "flag".
		name := strings.TrimLeft(arg, "-")
		// Skip non-flag tokens (positional args, flag values like "go" after --lang).
		if !strings.HasPrefix(arg, "-") || name == "" {
			continue
		}
		// For --flag=value style, only count the flag name part.
		name = strings.SplitN(name, "=", 2)[0]
		seen[name]++
		if seen[name] > 1 {
			return fmt.Errorf(ui.Error(fmt.Sprintf(messages.DuplicateFlag, arg)))
		}
	}
	return nil
}
