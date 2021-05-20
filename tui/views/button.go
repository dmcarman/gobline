package views

import (
	"gobline/tui/themes"

	tea "github.com/charmbracelet/bubbletea"
)

type ButtonView struct {
	selected bool
	disabled bool
	text     string
	on       func() tea.Msg
	theme    themes.Theme
}
type ButtonSelectMsg bool
type ButtonDisableMsg bool
type ButtonActivateMsg struct{}
type ButtonTextMsg string

func NewButtonView(selected bool, text string, on func() tea.Msg) *ButtonView {
	return &ButtonView{
		selected,
		false,
		text,
		on,
		themes.NewDefaultTheme(),
	}
}
func (v *ButtonView) WithStyle(theme themes.Theme) *ButtonView {
	v.theme = theme
	return v
}
func (v *ButtonView) Init() tea.Cmd {
	return nil
}
func (v *ButtonView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ButtonSelectMsg:
		v.selected = bool(msg)
		return v, nil
	case ButtonDisableMsg:
		v.disabled = bool(msg)
		return v, nil
	case ButtonTextMsg:
		v.text = string(msg)
		return v, nil
	case ButtonActivateMsg:
		if v.on != nil && !v.disabled {
			return v, v.on
		}
		return v, nil

	default:
		return v, nil
	}
}

func (v *ButtonView) View() string {
	if v.disabled {
		return v.theme.ButtonDisabled.Render(v.text)
	}
	if v.selected {
		return v.theme.ButtonHighlight.Render(v.text)
	}
	return v.theme.ButtonNormal.Render(v.text)
}
