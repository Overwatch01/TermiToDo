package tui

import "github.com/charmbracelet/lipgloss"

var homeLayoutStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1)

func RenderHome(m *Model) string {
	return homeLayoutStyle.Width(m.Width / 2).Render("This is the home layout")
}
