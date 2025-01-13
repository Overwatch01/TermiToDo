package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var menuBody string

type Model struct {
	Width      int
	Height     int
	CurrentTab int
	TaskInput  textinput.Model
}

type TermiKeys interface {
	SetKeys()
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// update the terminal dimension
		if msg.Width-20 > 60 {
			m.Width = msg.Width - 20
		}
		m.Height = 1000
	case tea.KeyMsg:
		if m.GetCurrentTab() == "task" {
			cmd = m.SetKeys(msg.String())
		}
		switch msg.String() {
		case "ctrl+c", "esc", "q":
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
		return m, cmd
	}
	return m, cmd
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

	textInputModel := textinput.New()
	textInputModel.Placeholder = "What do you plan to do today?"
	textInputModel.Focus()
	textInputModel.CharLimit = 250

	return Model{
		TaskInput: textInputModel,
	}
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
