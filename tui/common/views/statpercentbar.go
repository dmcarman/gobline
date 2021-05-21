package views

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatPercentbar struct {
	progress *progress.Model
	stat     *Stat
	Label    string
	color    string
}
type StatPercentBarUpdateMsg int
type StatPercentbarOption func(*StatPercentbar) error

func NewStatPercentbar(stat *Stat, label string, opts ...StatPercentbarOption) (*StatPercentbar, error) {
	// prog, err := progress.NewModel(progress.WithSolidFill(baseColor), progress.WithWidth(width))
	prog, err := progress.NewModel(progress.WithoutPercentage())
	if err != nil {
		return nil, err
	}
	prog.ShowPercentage = false
	statPercentBar := &StatPercentbar{
		prog,
		stat,
		label,
		""}
	for _, opt := range opts {
		if err := opt(statPercentBar); err != nil {
			return nil, err
		}
	}
	return statPercentBar, nil
}

func WithProgressWidth(width int) StatPercentbarOption {
	return func(sp *StatPercentbar) error {
		sp.progress.Width = width
		return nil
	}
}
func WithProgressSolidFill(color string) StatPercentbarOption {
	return func(sp *StatPercentbar) error {
		sp.progress.FullColor = color
		return nil
	}
}
func WithStatLabelColor(color string) StatPercentbarOption {
	return func(sp *StatPercentbar) error {
		sp.color = color
		return nil
	}
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
	case StatPercentBarUpdateMsg:
		f.stat.Update(StatUpdateValueMsg(int(msg)))
		return f, nil
	default:
		return f, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (f StatPercentbar) View() string {
	label := lipgloss.NewStyle().Foreground(lipgloss.Color(f.color)).Render(f.Label)
	progress := f.progress.View(float64(f.stat.value) / float64(f.stat.max))
	stat := lipgloss.PlaceHorizontal(6, lipgloss.Right, f.stat.View())
	return lipgloss.JoinHorizontal(lipgloss.Left, label, progress, stat)
}
