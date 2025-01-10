package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var items = []string{
	"The Romans learned from the Greeks that quinces slowly cooked with honey would “set” when cool. The Apicius gives a recipe for preserving whole quinces, stems and leaves attached, in a bath of honey diluted with defrutum: Roman marmalade. Preserves of quince and lemon appear (along with rose, apple, plum and pear) in the Book of ceremonies of the Byzantine Emperor Constantine VII Porphyrogennetos.",
	"The Romans learned from the Greeks that quinces slowly cooked with honey would “set” when cool. The Apicius gives a recipe for preserving whole quinces, stems and leaves attached, in a bath of honey diluted with defrutum: Roman marmalade. Preserves of quince and lemon appear (along with rose, apple, plum and pear) in the Book of ceremonies of the Byzantine Emperor Constantine VII Porphyrogennetos.",
	"Medieval quince preserves, which went by the French name cotignac, produced in a clear version and a fruit pulp version, began to lose their medieval seasoning of spices in the 16th century. In the 17th century, La Varenne provided recipes for both thick and clear cotignac.",
	"In 1524, Henry VIII, King of England, received a “box of marmalade” from Mr. Hull of Exeter. This was probably marmelada, a solid quince paste from Portugal, still made and sold in southern Europe today. It became a favourite treat of Anne Boleyn and her ladies in waiting.",
}

var homeLayoutStyle = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Border(lipgloss.DoubleBorder()).
	Foreground(lipgloss.Color("#FAFAFA")).
	Margin(1, 3, 0, 0).
	Padding(1, 2).
	Width(30)

func RenderHome(m *Model) string {
	itemWidth := homeLayoutStyle.GetWidth() + homeLayoutStyle.GetPaddingLeft() + homeLayoutStyle.GetPaddingRight()
	columns := m.Width / itemWidth

	if columns == 0 {
		columns = 1 //Ensure at least one column
	}

	var rows []string
	for i := 0; i < len(items); i += columns {
		end := i + columns
		if end > len(items) {
			end = len(items)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Left, renderItems(items[i:end])...)
		rows = append(rows, row)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func renderItems(items []string) []string {
	var rendered []string
	for _, item := range items {
		rendered = append(rendered, homeLayoutStyle.Align(lipgloss.Left).Render(item))
	}

	return rendered
}
