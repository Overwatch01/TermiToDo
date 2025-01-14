package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabModel struct {
	tab         string
	displayItem string
}

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

	tabItems = [...]TabModel{
		{
			tab:         "home",
			displayItem: "(H)ome",
		},
		{
			tab:         "task",
			displayItem: "(T)ask",
		},
		{
			tab:         "about",
			displayItem: "(A)bout",
		},
		{
			tab:         "help",
			displayItem: "(H)elp",
		},
	}
)

func RenderTab(m *Model) string {
	renderedTab := make([]string, len(tabItems))
	for i, item := range tabItems {
		if i == m.CurrentTab {
			renderedTab[i] = activeTab.Render(item.displayItem)
		} else {
			renderedTab[i] = tab.Render(item.displayItem)
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

func (m *Model) GetCurrentTab() TabModel {
	return tabItems[m.CurrentTab]

	// return strings.ToLower(tabItems[m.CurrentTab].tab)
}

func (m *Model) GetCurrentTabKeyMap(msg string) tea.Cmd {
	var cmd tea.Cmd
	cmd = m.DefaultKeyMap(msg)
	if cmd != nil {
		return cmd
	}
	tab := tabItems[m.CurrentTab].tab
	switch tab {
	case "task":
		task := TaskModel{m: m}
		cmd = task.SetKeyMap(msg)

	}
	return cmd
}
