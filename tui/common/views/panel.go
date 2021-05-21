package views

import (
	"gobline/tui/common/themes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Panel struct {
	Title   string
	content tea.Model
}

func NewPanel(title string, content tea.Model) *Panel {
	return &Panel{
		Title:   title,
		content: content,
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (p Panel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (p Panel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.content, cmd = p.content.Update(msg)
	return p, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (p Panel) View() string {

	panel := themes.DefaultPanelBorderBottom.Render(p.content.View())
	title := p.Title
	width := lipgloss.Width(panel)
	titleborder := lipgloss.PlaceHorizontal(width-2, lipgloss.Center, title)
	titleborder = themes.DefaultPanelTitleBorder.Render(titleborder)
	return lipgloss.JoinVertical(lipgloss.Left, titleborder, panel)

}
