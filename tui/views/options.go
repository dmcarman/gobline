package views

import (
	"gobline/tui/themes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OptionsView struct {
	selected int
	options  []ButtonAction
	vertical bool
	buttons  []tea.Model
	theme    themes.Theme
}

type ButtonAction struct {
	Label  string
	Action func() tea.Msg
}
type OptionsSelectMsg int

type OptionsActivateMsg struct{}

func newOptionsView(vertical bool, selected int, theme themes.Theme, options []ButtonAction) *OptionsView {
	btns := []tea.Model{}
	for i, t := range options {
		if i+1 == selected {
			btns = append(btns, NewButtonView(true, t.Label, t.Action).WithStyle(theme))
		} else {
			btns = append(btns, NewButtonView(false, t.Label, t.Action).WithStyle(theme))
		}
	}
	return &OptionsView{
		selected,
		options,
		vertical,
		btns,
		theme,
	}
}
func NewHorizontalOptionsView(selected int, options []ButtonAction) *OptionsView {
	theme := themes.NewDefaultTheme()
	return newOptionsView(false, selected, theme, options)
}
func NewVerticalOptionsView(selected int, options []ButtonAction) *OptionsView {
	theme := themes.NewDefaultTheme()
	return newOptionsView(true, selected, theme, options)
}
func (v *OptionsView) WithStyle(theme themes.Theme) *OptionsView {
	return newOptionsView(v.vertical, v.selected, theme, v.options)
}
func (v OptionsView) Init() tea.Cmd {
	return nil
}
func (v OptionsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case OptionsSelectMsg:
		return v.handleSelect(int(msg))
	case OptionsActivateMsg:
		button, cmd := v.buttons[v.selected].Update(ButtonActivateMsg{})
		v.buttons[v.selected] = button
		return v, cmd
	default:
		return v, nil
	}
}

func (v OptionsView) View() string {
	var s = ""
	if v.vertical {
		for _, t := range v.buttons {
			s = lipgloss.JoinVertical(lipgloss.Left, s, t.View())
		}
	} else {
		for _, t := range v.buttons {
			s = lipgloss.JoinHorizontal(lipgloss.Bottom, s, t.View())
		}
	}
	return s
}

func (v OptionsView) handleSelect(selected int) (tea.Model, tea.Cmd) {
	for i := range v.buttons {
		if i+1 == selected {
			v.buttons[i].Update(ButtonSelectMsg(true))
		} else {
			v.buttons[i].Update(ButtonSelectMsg(false))
		}
	}
	return v, nil
}
