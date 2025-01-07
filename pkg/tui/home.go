package tui

import "github.com/charmbracelet/lipgloss"

func GetHomeStyle() string {
	return lipgloss.NewStyle().Render("This is the home layout")

}
