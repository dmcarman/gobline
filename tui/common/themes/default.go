package themes

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	ButtonNormal    lipgloss.Style
	ButtonHighlight lipgloss.Style
	ButtonDisabled  lipgloss.Style
}

func NewDefaultTheme() Theme {
	return Theme{
		DefaultButtonNormal,
		DefaultButtonHighlight,
		DefaultButtonDisabled}
}

var DefaultButtonNormal lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Padding(0, 2).
	Margin(0, 1)
var DefaultButtonHighlight lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#1D0660")).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#1D0660")).
	Padding(0, 2).
	Margin(0, 1)
var DefaultButtonDisabled lipgloss.Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#063260")).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#063260")).
	Padding(0, 2).
	Margin(0, 1)
var DefaultPanelBorderBottom lipgloss.Style = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), false, true, true, true)
var DefaultPanelTitleBorder lipgloss.Style = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true, true, false, true)
var TotalTheme lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("#FAFAFA")).Foreground(lipgloss.Color("#0F0F0F")).PaddingRight(1).MarginLeft(1)
var ItemlistNameTheme lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4"))
var ItemlistDescriptionTheme lipgloss.Style = lipgloss.NewStyle().BorderLeft(true).BorderStyle(lipgloss.Border{
	Top:         "",
	Bottom:      "",
	Left:        "|",
	Right:       "",
	TopLeft:     "",
	TopRight:    "",
	BottomRight: "",
	BottomLeft:  "",
})
