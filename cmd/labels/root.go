package labels

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "labels",
	Short: "A brief description of your command",
}

func init() {
	RootCommand.AddCommand(ListCommand)
}
