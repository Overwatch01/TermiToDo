package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var menuBody string

type Model struct {
	Width      int
	Height     int
	CurrentTab int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// update the terminal dimension
		m.Width = msg.Width - 20
		m.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "k":
			if m.CurrentTab > 0 {
				m.CurrentTab--
			}
		case "down", "j":
			if m.CurrentTab < m.GetMenuCount() {
				m.CurrentTab++
			}
		}

		menuBody = m.getMenuBody()
	default:
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	layoutStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.Width).
		Padding(1, 2).
		Margin(1, 10)

	header := RenderHeader(&m)
	menu := RenderMenu(&m)
	body := lipgloss.JoinHorizontal(lipgloss.Top, menu, menuBody)
	layout := lipgloss.JoinVertical(lipgloss.Top, header, body)
	return layoutStyle.Render(layout)
}

func InitialModel() Model {
	return Model{}
}

func (m *Model) getMenuBody() string {
	switch m.GetCurrentMenu() {
	case "help":
		return RenderHelp()

	default:
		return RenderHome(m)
	}
}
