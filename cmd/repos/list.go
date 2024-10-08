package repos

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/google/go-github/v65/github"
	"github.com/spf13/cobra"
	"gru/pkg/console"
	gh "gru/pkg/github"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List all existing repos for the selected GitHub owner (user, org)",
	Run: func(cmd *cobra.Command, args []string) {
		config := gh.BuildGitHubConfig()
		opt := &github.RepositoryListByOrgOptions{}
		repos, _, err := gh.Client().Repositories.ListByOrg(
			context.Background(),
			gh.BuildGitHubConfig().Organization,
			opt,
		)

		if err != nil {
			log.Fatalf("Error listing repositories: %v", err)
		}

		log.Infof("List of all repositories for org: %s", config.Organization)
		console.ReposTable(
			console.ReposTableData{
				Repositories: repos,
			},
		)

	},
}
