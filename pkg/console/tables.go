package console

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/google/go-github/v65/github"
	"time"
)

type ReposTableData struct {
	Repositories []*github.Repository
}

func ReposTable(data ReposTableData) {
	t := table.New()
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99"))
	headerStyle := lipgloss.NewStyle().PaddingLeft(2).PaddingRight(2).Bold(true)
	rowStyle := lipgloss.NewStyle().PaddingLeft(2).PaddingRight(2)

	t.Border(lipgloss.NormalBorder())
	t.Width(80)
	t.StyleFunc(func(row, col int) lipgloss.Style {
		switch {
		case row == 0:
			return headerStyle
		default:
			return rowStyle
		}
	})
	t.BorderStyle(borderStyle)
	t.Headers(
		"NAME",
		"VISIBILITY",
		"LAST UPDATED",
	)

	for _, repo := range data.Repositories {
		t.Row(
			repo.GetName(),
			repo.GetVisibility(),
			repo.GetUpdatedAt().Format(time.RFC1123Z),
		)
	}

	fmt.Println(t)
}
