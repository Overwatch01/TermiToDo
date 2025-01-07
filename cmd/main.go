package main

import (
	"log"

	"github.com/Overwatch01/TermToDo/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Unable to run bubble tea: %s", err)
	}
}
