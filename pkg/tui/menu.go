package tui

import "github.com/charmbracelet/lipgloss"

var menuStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2)

var selectedMenuStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("73")).
	Padding(1, 1).
	Margin(0)

var unselectedMenuStyle = lipgloss.NewStyle().
	Padding(0, 1).
	Margin(0)

var menus = [...]string{"Home", "Task", "About", "Auth", "Help"}

func RenderMenu(m *Model) string {
	var menuItems string
	width := m.Width / 4
	for i, menu := range menus {

		if i == m.CurrentTab {
			menuItems += selectedMenuStyle.Width(width).Render(" "+menu) + "\n"
		} else {
			menuItems += unselectedMenuStyle.Render("> "+menu) + "\n"
		}
	}

	return menuStyle.Render(menuItems)
}

func (m *Model) GetMenuCount() int {
	return len(menus) - 1
}
