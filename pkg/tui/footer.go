package tui

import "github.com/charmbracelet/lipgloss"

func RenderFooter() string {
	return lipgloss.NewStyle().Render("This is the footer page where we would have all the keys you can use for navigation")
}
