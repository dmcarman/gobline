package views

import (
	"fmt"
	"gobline/tui/common/themes"

	tea "github.com/charmbracelet/bubbletea"
)

type Stat struct {
	value     int
	max       int
	formatter func(int, int) string
}
type StatUpdateMaxMsg int
type StatUpdateValueMsg int
type StatOption func(f *Stat) error

func NewStat(value, max int, opts ...StatOption) (*Stat, error) {
	defaultFormatter := func(value int, max int) string { return fmt.Sprintf("%v/%v", value, max) }
	stat := &Stat{
		value:     value,
		max:       max,
		formatter: defaultFormatter,
	}
	for _, opt := range opts {
		if err := opt(stat); err != nil {
			return nil, err
		}
	}
	stat.limitStatBounds()
	return stat, nil
}
func WithStatFormatter(formatter func(int, int) string) StatOption {
	return func(f *Stat) error {
		f.formatter = formatter
		return nil
	}
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
	return themes.TotalTheme.Render(f.formatter(f.value, f.max))
}

func (f *Stat) limitStatBounds() {
	if f.value > f.max {
		f.value = f.max
	} else if f.value < 0 {
		f.value = 0
	}
}
