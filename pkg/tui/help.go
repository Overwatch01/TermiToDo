package tui

import "github.com/charmbracelet/lipgloss"

func RenderHelpStyle() string {
	return lipgloss.NewStyle().Render("This is the help layout`")
}
