package tags

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "tags",
	Short: "Manage tags",
}

func init() {
	RootCommand.AddCommand(ListCommand)
}
