package github

import (
	"github.com/google/go-github/v65/github"
	"gru/pkg/utils"
)

type Config struct {
	Organization string
	Repository   string
	Token        string
}

func Client() *github.Client {
	config := BuildGitHubConfig()
	client := github.NewClient(nil).WithAuthToken(config.Token)
	return client
}

func BuildGitHubConfig() Config {
	org := utils.GetEnvironment("GITHUB_OWNER")
	repo := utils.GetEnvironment("GITHUB_REPO")
	token := utils.GetEnvironment("GITHUB_TOKEN")

	return Config{
		Organization: org,
		Repository:   repo,
		Token:        token,
	}
}
