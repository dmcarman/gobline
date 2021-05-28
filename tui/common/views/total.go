package views

import (
	"fmt"
	"gobline/tui/common/themes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Total struct {
	Label          string
	Amount         int
	maxLabelWidth  int
	maxAmountWidth int
	formatString   string
}
type TotalAmountUpdateMsg int
type TotalOption func(*Total) error

func NewTotal(label string, amount int, opts ...TotalOption) (*Total, error) {
	total := &Total{
		Label:          label,
		Amount:         amount,
		maxLabelWidth:  0,
		maxAmountWidth: 0,
		formatString:   "%d",
	}
	for _, opt := range opts {
		if err := opt(total); err != nil {
			return nil, err
		}
	}

	return total, nil
}
func WithMaxLabelWidth(width int) TotalOption {
	return func(s *Total) error {
		s.maxLabelWidth = width
		return nil
	}
}
func WithMaxAmountWidth(width int) TotalOption {
	return func(s *Total) error {
		s.maxAmountWidth = width
		return nil
	}
}
func WithColor(color string) TotalOption {
	return func(s *Total) error {
		s.Label = lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(s.Label)
		return nil
	}
}
func WithTotalFormatString(formatString string) TotalOption {
	return func(s *Total) error {
		s.formatString = formatString
		return nil
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s Total) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s Total) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TotalAmountUpdateMsg:
		s.Amount = int(msg)
		return s, nil
	default:
		return s, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s Total) View() string {
	var label string
	var amount string
	if s.maxLabelWidth > 0 {
		label = lipgloss.PlaceHorizontal(s.maxLabelWidth, lipgloss.Left, s.Label, lipgloss.WithWhitespaceChars("."))
	} else {
		label = s.Label
	}
	if s.maxAmountWidth > 0 {
		amount = lipgloss.PlaceHorizontal(s.maxAmountWidth, lipgloss.Right, fmt.Sprintf(s.formatString, s.Amount), lipgloss.WithWhitespaceChars(" "))
	} else {
		amount = fmt.Sprintf(":%v", s.Amount)
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, label, themes.TotalTheme.Render(amount))
}
