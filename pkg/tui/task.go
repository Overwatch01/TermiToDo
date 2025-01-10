package tui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	list = lipgloss.NewStyle()

	listHeader = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EEEEEE")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderLeft(true).
			Padding(0, 1)

	taskInfo = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder())

	// taskItems []string = { "Item 1", "Item 2"}
)

func RenderTask(m *Model) string {

	width := m.Width - 10
	taskInfoWidth := width / 4
	taskWidth := width - taskInfoWidth
	tasks := lipgloss.JoinVertical(lipgloss.Left, renderTaskItem(taskWidth-10)...)
	// tasks := lipgloss.JoinVertical(lipgloss.Left, listHeader.Width(taskWidth-10).Render("Task Completed 1/10"), listHeader.Render("Item no 1"), listHeader.Render("Item no 2"), listHeader.Render("Item no 3"), listHeader.Render("Item no 4"))

	widthInfo := fmt.Sprintf("\n Width: %v \n taskInfoWidth: %v \n taskWidth: %v", width, taskInfoWidth, width-taskInfoWidth)
	info := taskInfo.Width(taskInfoWidth).Render("This will show the information on the task that needs to be completed " + widthInfo)
	return lipgloss.JoinHorizontal(lipgloss.Left, info, list.Width(taskWidth).Render(tasks))
}

func renderTaskItem(width int) []string {
	var taskItems []string

	for i := 0; i <= 10; i++ {
		taskItems = append(taskItems, listHeader.Width(width).Render("Item no "+strconv.Itoa(i)))
	}

	return taskItems
}
