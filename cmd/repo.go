package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"%s: %s",
			"Current operating system:", "me",
		)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
