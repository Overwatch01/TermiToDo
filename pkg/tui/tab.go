package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	tabItems = [...]string{"Home", "Task", "About", "Auth", "Help"}
)

func RenderTab(m *Model) string {
	renderedTab := make([]string, len(tabItems))
	for i, item := range tabItems {
		if i == m.CurrentTab {
			renderedTab[i] = activeTab.Render(item)
		} else {
			renderedTab[i] = tab.Render(item)
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTab...)
	m.Width = max(0, m.Width-lipgloss.Width(row)-2)
	gap := tabGap.Render(strings.Repeat(" ", m.Width))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	return row
}

func (m *Model) GetTabCount() int {
	return len(tabItems) - 1

}

func (m *Model) GetCurrentTab() string {
	return strings.ToLower(tabItems[m.CurrentTab])
}
