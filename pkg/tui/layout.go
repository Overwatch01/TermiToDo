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
		if msg.Width-20 > 60 {
			m.Width = msg.Width - 20
		}
		m.Height = 1000
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "left", "h":
			if m.CurrentTab > 0 {
				m.CurrentTab--
			}
		case "right", "l":
			if m.CurrentTab < m.GetTabCount() {
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
		Border(lipgloss.RoundedBorder(), false, true, true, true).
		Width(m.Width).
		Margin(1, 10)

	tab := RenderTab(&m)
	// header := RenderHeader(&m)
	// menu := RenderMenu(&m)
	body := lipgloss.JoinVertical(lipgloss.Top, tab, menuBody)
	// layout := lipgloss.JoinVertical(lipgloss.Top, header, body)
	return layoutStyle.Render(body)
}

func InitialModel() Model {
	return Model{}
}

func (m Model) getMenuBody() string {
	switch m.GetCurrentTab() {
	case "help":
		return RenderHelp()

	case "task":
		return RenderTask(&m)

	default:
		return RenderHome(&m)
	}
}
