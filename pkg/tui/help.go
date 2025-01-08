package tui

import "github.com/charmbracelet/lipgloss"

func RenderHelp() string {
	return lipgloss.NewStyle().Render("This is the help layout`")
}
