package tui

import "github.com/charmbracelet/lipgloss"

func RenderAbout() string {
	return lipgloss.NewStyle().Render("This is the about layout`")
}
