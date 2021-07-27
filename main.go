package main

import (
	"fmt"
	"os"

	"github.com/dmcarman/gobline/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model, err := tui.NewTUI()
	if err != nil {
		fmt.Println("Could not intialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if err := tea.NewProgram(model).Start(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}
