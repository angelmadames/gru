package cmd

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/google/go-github/v65/github"
	"github.com/spf13/cobra"
	gh "gru/pkg/github"
)

var repoCmd = &cobra.Command{
	Use:   "repos",
	Short: "Manage repositories",
	Run: func(cmd *cobra.Command, args []string) {
		opt := &github.RepositoryListByOrgOptions{}
		repos, _, err := gh.Client().Repositories.ListByOrg(
			context.Background(),
			gh.BuildGitHubConfig().Organization,
			opt,
		)
		if err != nil {
			log.Fatalf("Error listing repositories: %v", err)
		}

		for _, repo := range repos {
			log.Info("GitHub repository info:", "repo", repo.GetName(), "desc", repo.GetDescription())
		}
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
