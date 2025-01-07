package tui

import (
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

func getHeaderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Width(120).
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("63")).
		Padding(1, 2)

}

func RenderHeader(m *Model) string {
	return getHeaderStyle().Render("This is the header with a width of  just to also test the width I am putting some lorem ipsum characters" + strconv.Itoa(m.Width))
}
