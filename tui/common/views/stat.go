package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Stat struct {
	value int
	max   int
}
type StatUpdateMaxMsg int
type StatUpdateValueMsg int

func NewStat(value, max int) *Stat {
	stat := &Stat{
		value: value,
		max:   max,
	}
	stat.limitStatBounds()
	return stat
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (f Stat) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (f *Stat) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatUpdateMaxMsg:
		f.max = int(msg)
		f.limitStatBounds()
		return f, nil
	case StatUpdateValueMsg:
		f.value = int(msg)
		f.limitStatBounds()
		return f, nil
	default:
		return f, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (f Stat) View() string {
	return fmt.Sprintf("[%+2d] %2d", (f.value-10)/2, f.value)
}

func (f *Stat) limitStatBounds() {
	if f.value > f.max {
		f.value = f.max
	} else if f.value < 0 {
		f.value = 0
	}
}
