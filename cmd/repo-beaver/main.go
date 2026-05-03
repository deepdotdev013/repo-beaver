package main

import (
	"fmt"
	"time"

	"github.com/deepdotdev013/repo-beaver/internal/contracts"
	"github.com/deepdotdev013/repo-beaver/internal/generator"
	"github.com/deepdotdev013/repo-beaver/internal/policy"
	"github.com/deepdotdev013/repo-beaver/internal/prompt"
	"github.com/deepdotdev013/repo-beaver/pkg/errors"
	"github.com/deepdotdev013/repo-beaver/pkg/messages"
	"github.com/deepdotdev013/repo-beaver/pkg/ui"
)

func main() {
	fmt.Println(ui.Primary("🦫 Repo Beaver"))
	fmt.Println(ui.White("Generate backend projects in seconds.\n"))
	// Step 1: Start the project name & language selection prompt.
	projectName, language, modulePath, framework, err := prompt.StartLanguagePrompt()
	errors.HandleError(err, 1)

	// Step 2: Check for language-specific dependencies.
	policyErr := policy.CheckLanguageDeps(language)
	if policyErr != nil {
		errors.HandleError(policyErr, 1)
	}

	// Step 3: Check for directory overwrite policy.
	ok, err := policy.AvoidDirOverwrite(projectName)
	errors.HandleError(err, 1)
	if !ok {
		fmt.Println(messages.OverwriteCancelled)
		return
	}

	// Step 4: Get the appropriate generator based on the selected language.
	gen, err := generator.Get(language)
	errors.HandleError(err, 1)

	// Record the start time for performance measurement.
	start := time.Now()

	// Step 5: Generate the project structure.
	err = ui.RunSpinner(ui.Success(messages.CreatingProjectStructure), func() error {
		return gen.Generate(contracts.InitConfig{
			ProjectName: projectName,
			ModulePath:  modulePath,
			Framework:   framework,
		})
	})
	errors.HandleError(err, 1)

	// Step 6: Initialize the project with necessary configurations.
	err = ui.RunSpinner(ui.Success(messages.InitializingProject), func() error {
		return gen.Init(contracts.InitConfig{
			ProjectName: projectName,
			ModulePath:  modulePath,
			Framework:   framework,
		})
	})
	errors.HandleError(err, 1)

	duration := time.Since(start)
	fmt.Println(ui.Success(fmt.Sprintf(messages.ProjectGeneratedSuccess, duration.Seconds())))
}
