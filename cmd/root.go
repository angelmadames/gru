package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ops",
	Short: "An all-purpose deployment automation tool tailored for DevOps & SRE",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP(
		"toggle",
		"t",
		false,
		"Help message for toggle",
	)
}
