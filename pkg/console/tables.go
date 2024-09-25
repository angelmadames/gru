package console

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/google/go-github/v65/github"
)

type ReposTableData struct {
	Repositories []*github.Repository
}

func ReposTable(data ReposTableData) {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("20"))).
		Headers("NAME", "DEFAULT BRANCH")

	for _, repo := range data.Repositories {
		t.Row(repo.GetName(), repo.GetDefaultBranch())
	}

	fmt.Println(t)
}
