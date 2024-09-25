package cmd

import (
	"gru/cmd/labels"
	"gru/cmd/tags"
	"os"

	"github.com/spf13/cobra"
	"gru/cmd/repos"
)

var rootCmd = &cobra.Command{
	Use:   "gru",
	Short: "CLI tool to run maintenance tasks on GitHub repositories",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(repos.RootCommand)
	rootCmd.AddCommand(labels.RootCommand)
	rootCmd.AddCommand(tags.RootCommand)
}
