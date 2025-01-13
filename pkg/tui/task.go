package tui

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	taskItems []string = make([]string, 0)

	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	list = lipgloss.NewStyle()

	listHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EEEEEE")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderBottom(true).
			Padding(0, 1)

	taskInfo = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder())

	taskInput = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderBottom(true)

	userInput string
	// taskItems []string = { "Item 1", "Item 2"}
)

func RenderTask(m *Model) string {
	width := m.Width - 10
	taskInfoWidth := width / 4
	taskWidth := width - taskInfoWidth - 10
	tasks := lipgloss.JoinVertical(lipgloss.Left, renderTaskItem(taskWidth)...)

	widthInfo := fmt.Sprintf("\n Width: %v \n taskInfoWidth: %v \n taskWidth: %v \n User Input: %v", width, taskInfoWidth, width-taskInfoWidth, m.TaskInput.Value())

	info := taskInfo.Width(taskInfoWidth).Render("This will show the information on the task that needs to be completed " + widthInfo)
	m.TaskInput.Width = taskWidth
	taskView := lipgloss.JoinVertical(lipgloss.Left, taskInput.Width(taskWidth).Render(m.TaskInput.View()), list.Width(taskWidth).Render(tasks))

	return lipgloss.JoinHorizontal(lipgloss.Left, info, taskView)
}

func renderTaskItem(width int) []string {
	if len(taskItems) > 0 {
		return taskItems
	}
	for i := 0; i <= len(taskItems)-1; i++ {
		taskItems = append(taskItems, listHeader.Width(width).Render("Item no "+strconv.Itoa(i)))
	}
	return taskItems
}

func (m *Model) SetKeys(msg string) tea.Cmd {
	var cmd tea.Cmd
	switch msg {
	case "enter":
		if m.TaskInput.Value() != "" {
			taskItems = append(taskItems, m.TaskInput.Value())
			userInput = ""
		}
	case "backspace":
		if m.TaskInput.Value() != "" {
			// userInput = strings.TrimSuffix(userInput)
			userInput = userInput[:len(userInput)-1]
		}
	case "tab":
		break

	default:
		userInput = userInput + msg
	}
	m.TaskInput.SetValue(userInput)
	m.TaskInput, cmd = m.TaskInput.Update(msg)
	return cmd
}
