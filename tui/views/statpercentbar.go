package views

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatPercentbar struct {
	progress *progress.Model
	percent  float64
	Label    string
}
type StatUpdateMsg float64

func NewStatPercentbar(percent float64, width int, label string) (*StatPercentbar, error) {
	// prog, err := progress.NewModel(progress.WithScaledGradient(colorFrom, colorTo), progress.WithWidth(width))
	prog, err := progress.NewModel(progress.WithWidth(width))
	if err != nil {
		return nil, err
	}
	prog.ShowPercentage = false
	return &StatPercentbar{
		prog,
		percent,
		label}, nil
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (f StatPercentbar) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (f StatPercentbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatUpdateMsg:
		f.percent = float64(msg)
		return f, nil
	default:
		return f, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (f StatPercentbar) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, f.Label, " ", f.progress.View(f.percent)+" ")
}
