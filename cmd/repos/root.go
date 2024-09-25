package repos

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "repos",
	Short: "Manage repositories",
}

func init() {
	RootCommand.AddCommand(ListCommand)
}
