package cmd

import (
	"github.com/charmbracelet/log"

	"github.com/spf13/cobra"
)

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"%s: %s",
			"Current operating system:", "me",
		)
	},
}

func init() {
	rootCmd.AddCommand(labelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
