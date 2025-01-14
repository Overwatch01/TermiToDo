package tui

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TaskModel struct {
	m *Model
}

func (tm TaskModel) SetKeyMap(msg string) tea.Cmd {
	var cmd tea.Cmd
	switch msg {
	case "enter":
		if tm.m.TaskInput.Value() != "" {
			taskItems = append(taskItems, tm.m.TaskInput.Value())
			tm.m.InputMode = false
			userInput = ""
		}
	case "backspace":
		if tm.m.TaskInput.Value() != "" {
			userInput = userInput[:len(userInput)-1]
		}
	case "tab", "right", "left":
		break
	case "esc":
		tm.m.InputMode = false
	case "up":
		if selectedTask > 0 {
			tm.m.InputMode = false
			selectedTask--
		}
	case "down":
		if selectedTask < len(taskItems)-1 {
			tm.m.InputMode = false
			selectedTask++
		}
	default:
		tm.m.InputMode = true
		userInput = userInput + msg
	}
	tm.m.TaskInput.SetValue(userInput)
	tm.m.TaskInput, cmd = tm.m.TaskInput.Update(msg)
	return cmd
}

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	list = lipgloss.NewStyle()

	focusedTaskList = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EEEEEE")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(highlight).
			Padding(0, 1)

	unfocusedTaskList = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#EEEEEE")).
				BorderStyle(lipgloss.NormalBorder()).
				BorderLeft(true).
				BorderBottom(true).
				Padding(0, 1)

	taskInfo = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder())

	unfocusedTaskInput = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderBottom(true)

	focusedTaskInput = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(highlight)

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	selectedTask = 0

	userInput string
	taskItems []string = []string{"Item 1", "Item 2"}
)

func RenderTask(m *Model) string {
	var taskInput lipgloss.Style

	width := m.Width - 10
	taskInfoWidth := width / 4
	taskWidth := width - taskInfoWidth - 10
	tasks := lipgloss.JoinVertical(lipgloss.Left, renderTaskItem(taskWidth)...)

	widthInfo := fmt.Sprintf("\n Width: %v \n taskInfoWidth: %v \n taskWidth: %v \n User Input: %v \n Number of Tasks: %v \n Input Mode: %v", width, taskInfoWidth, width-taskInfoWidth, m.TaskInput.Value(), len(taskItems), m.InputMode)

	info := taskInfo.Width(taskInfoWidth).Render("This will show the information on the task that needs to be completed " + widthInfo)
	m.TaskInput.Width = taskWidth

	if m.InputMode {
		taskInput = focusedTaskInput
	} else {
		taskInput = unfocusedTaskInput
	}

	taskView := lipgloss.JoinVertical(lipgloss.Left, taskInput.Width(taskWidth).Render(m.TaskInput.View()), list.Render(tasks))

	taskUI := lipgloss.JoinHorizontal(lipgloss.Left, info, taskView)

	return lipgloss.NewStyle().Render(taskUI, renderDialog(m.Width))
}

func getTask() {
	if len(taskItems) < 0 {
		for i := 0; i < 10; i++ {
			taskItems = append(taskItems, "Item no "+strconv.Itoa(i))
		}
	}
}

func renderTaskItem(width int) []string {
	getTask()
	var taskList lipgloss.Style
	renderTask := []string{}
	for i, item := range taskItems {
		if i == selectedTask {
			taskList = focusedTaskList
		} else {
			taskList = unfocusedTaskList
		}
		renderTask = append(renderTask, taskList.Width(width).Render(item))
	}
	return renderTask

}

func renderDialog(width int) string {
	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Do you want to mark task as completed?")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	return dialog

}
