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
	InputMode  bool
	TaskMode   TaskMode
}

type KeyMap interface {
	SetKeyMap(msg string) tea.Cmd
}

func (m *Model) DefaultKeyMap(msg string) tea.Cmd {
	var cmd tea.Cmd = nil
	switch msg {
	case "ctrl+c":
		cmd = tea.Quit
	case "q":
		if !m.InputMode {
			cmd = tea.Quit
		}
	case "esc":
		if m.InputMode == true {
			m.InputMode = false
		} else {
			cmd = tea.Quit
		}
	case "left":
		if m.CurrentTab > 0 && !m.InputMode {
			m.CurrentTab--
		}
	case "right":
		if m.CurrentTab < m.GetTabCount() && !m.InputMode {
			m.CurrentTab++
		}

	}
	return cmd
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
		// m.SetKeys(layoutaLayout)
		cmd = m.GetCurrentTabKeyMap(msg.String())
		// if m.GetCurrentTab() == "task" {
		// 	cmd = m.SetKeys(msg.String())
		// }
		// switch msg.String() {
		// case "ctrl+c", "esc", "q":
		// 	return m, tea.Quit
		// case "left", "h":
		// 	if m.CurrentTab > 0 {
		// 		m.CurrentTab--
		// 	}
		// case "right", "l":
		// 	if m.CurrentTab < m.GetTabCount() {
		// 		m.CurrentTab++
		// 	}
		// }

	default:
		return m, cmd
	}
	menuBody = m.getMenuBody()
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
		InputMode: false,
		TaskMode:  All,
	}
}

func (m *Model) getMenuBody() string {
	switch m.GetCurrentTab().tab {
	case "help":
		return RenderHelp()

	case "task":
		return RenderTask(m)

	default:
		return RenderHome(m)
	}
}
