package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.12"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "repo-beaver",
	Short: "Generate backend projects in seconds",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Repo Beaver version:", version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
