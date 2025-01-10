package tui

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func getHeaderStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63")).
		Padding(1, 2)

}

func RenderHeader(m *Model) string {
	return getHeaderStyle(m.Width).Render("This is the header with a width of  just to also test the width I am putting some lorem ipsum characters" + strconv.Itoa(m.Width))
}
