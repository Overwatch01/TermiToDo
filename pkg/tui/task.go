package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/Overwatch01/TermToDo/model"
	"github.com/Overwatch01/TermToDo/pkg/file"
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
		if tm.m.TaskInput.Value() != "" && tm.m.InputMode {
			taskItems = NewTask(tm.m.TaskInput.Value())
			tm.m.InputMode = false
			userInput = ""
		}

		if !tm.m.InputMode && tm.m.TaskInput.Value() == "" {
			showDialog = true
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
		if strings.HasPrefix(msg, "ctrl") {
			if msg == "ctrl+u" {
				taskItems
			}
			break
		}

		if showDialog {
			if msg == "n" {
				showDialog = false
				userInput = ""
			}

			if msg == "y" {
				setTaskAsCompleted()
				showDialog = false
				userInput = ""
			}
		} else {
			tm.m.InputMode = true
			userInput = userInput + msg
		}
	}
	tm.m.TaskInput.SetValue(userInput)
	tm.m.TaskInput.CursorEnd()
	tm.m.TaskInput, cmd = tm.m.TaskInput.Update(msg)
	return cmd
}

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	list = lipgloss.NewStyle()

	focusedTaskList = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EEEEEE")).
			BorderStyle(lipgloss.DoubleBorder()).
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
	showDialog   = false

	userInput string
	taskItems = []model.Task{}
)

func RenderTask(m *Model) string {
	var taskInput lipgloss.Style
	var taskView string

	width := m.Width - 10
	taskInfoWidth := width / 4
	taskWidth := width - taskInfoWidth - 10
	tasks := lipgloss.JoinVertical(lipgloss.Left, renderTaskItem(taskWidth)...)

	widthInfo := fmt.Sprintf("\n Your task for %v \n Keep going 👊", time.Now().Format("2006-01-02"))
	// widthInfo := fmt.Sprintf("\n Width: %v \n taskInfoWidth: %v \n taskWidth: %v \n User Input: %v \n Number of Tasks: %v \n Input Mode: %v", width, taskInfoWidth, width-taskInfoWidth, m.TaskInput.Value(), len(taskItems), m.InputMode)

	info := taskInfo.Width(taskInfoWidth).Render(widthInfo + getTasksInfo())
	m.TaskInput.Width = taskWidth

	if m.InputMode {
		taskInput = focusedTaskInput
	} else {
		taskInput = unfocusedTaskInput
	}

	if !showDialog {
		taskView = lipgloss.JoinVertical(lipgloss.Left, taskInput.Width(taskWidth).Render(m.TaskInput.View()), list.Render(tasks))

	} else {
		m.TaskInput.SetValue(taskItems[selectedTask].Task)
		taskView = lipgloss.JoinVertical(lipgloss.Left, taskInput.Width(taskWidth).Render(m.TaskInput.View()), renderDialog(taskWidth))
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, info, taskView)

}

func getTask() {
	taskItems, _ = file.ReadFile()
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
		renderTask = append(renderTask, taskList.Width(width).Render(item.Task))
	}
	return renderTask

}

func renderDialog(width int) string {
	okButton := activeButtonStyle.Render("(Y)es")
	cancelButton := buttonStyle.Render("(N)o")

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Do you want to mark task as completed?")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 29,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	return dialog

}

func NewTask(task string) []model.Task {
	taskItems = append(taskItems, model.Task{Id: len(taskItems), Task: task, Completed: false})
	file.SaveFile(taskItems)
	return taskItems
}

func setTaskAsCompleted() {
	taskItems[selectedTask].Completed = true
	file.SaveFile(taskItems)
	return
}

func getTasksInfo() string {
	return fmt.Sprintf("\n\n Tasks (%v) \n Uncompleted Tasks (%v) \n Completed Tasks(%v)\n ", len(taskItems), getTotalTaskByStatus(false), getTotalTaskByStatus(true))
}

func getTotalTaskByStatus(completed bool) int {
	totalCount := 0

	for _, item := range taskItems {
		if item.Completed == completed {
			totalCount++
		}
	}
	return totalCount
}
