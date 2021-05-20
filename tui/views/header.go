package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HeaderView struct {
	Title string
	style lipgloss.Style
}

func NewHeaderView(text string) *HeaderView {
	return &HeaderView{
		text,
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Width(120).
			Align(lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FAFAFA")),
	}
}
func (v HeaderView) Init() tea.Cmd {
	return nil
}
func (v HeaderView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, nil
}

func (v HeaderView) View() string {
	return v.style.Render(v.Title)
}
