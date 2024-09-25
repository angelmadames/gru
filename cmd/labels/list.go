package labels

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "labels",
	Short: "Manage labels",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"%s: %s",
			"Current operating system:", "me",
		)
	},
}
