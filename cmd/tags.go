package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"%s: %s",
			"Current operating system:", "me",
		)
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
