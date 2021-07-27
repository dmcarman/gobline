package views_test

import (
	"gobline/tui/common/themes"
	"gobline/tui/common/views"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func getTestOptionView(selected int) tea.Model {
	testTheme := themes.NewDefaultTheme()
	testTheme.ButtonNormal = lipgloss.NewStyle()
	testTheme.ButtonHighlight = lipgloss.NewStyle().Bold(true)
	testTheme.ButtonDisabled = lipgloss.NewStyle().Faint(true)
	options := []views.ButtonAction{
		{"T", func() tea.Msg { return "Button 1" }},
		{"T", func() tea.Msg { return "Button 2" }},
		{"T", func() tea.Msg { return "Button 3" }},
	}
	optview := views.NewHorizontalOptionsView(selected, options).WithStyle(testTheme)
	return optview
}

func TestOptionsInitIsNil(t *testing.T) {
	optview := getTestOptionView(1)
	cmd := optview.Init()
	if cmd != nil {
		t.Fatalf("optionView init method should return nil")
	}
}
func TestOptionsNegativeSelectionHighlightsNone(t *testing.T) {
	optview := getTestOptionView(-1)
	if optview.View() != "TTT" {
		t.Fatal("Should not decorate any option with negative default")
	}
	optview.Update(views.OptionsSelectMsg(-1))
	if optview.View() != "TTT" {
		t.Fatal("Should not decorate any option with negative default")
	}

}
func TestOptionsSelectionExceedsRange(t *testing.T) {
	optview := getTestOptionView(99)
	if optview.View() != "TTT" {
		t.Fatal("Should not decorate any option when exceeding option range")
	}
	optview.Update(views.OptionsSelectMsg(99))
	if optview.View() != "TTT" {
		t.Fatal("Should not decorate any option when exceeding option range")
	}
}
func TestOptionsHighlightSelected(t *testing.T) {
	optview := getTestOptionView(2)
	if optview.View() != "T\x1B[1mT\x1B[0mT" {
		t.Log(optview.View())
		t.Fatal("Selected Option not highlighted")
	}
	optview.Update(views.OptionsSelectMsg(3))
	if optview.View() != "TT\x1B[1mT\x1B[0m" {
		t.Log(optview.View())
		t.Fatal("Selected Option not highlighted")
	}
}

func TestOptionsActivateSelected(t *testing.T) {
	optview := getTestOptionView(1)
	_, cmd := optview.Update(views.OptionsActivateMsg{})
	if cmd == nil {
		t.Fatalf("cmd not returned from Update method")
	}
}

func TestOptionsUnrecognizedMsg(t *testing.T) {
	optview := getTestOptionView(-1)
	_, cmd := optview.Update(testMsg{})
	if optview.View() != "TTT" {
		t.Fatal("Nothing should happen on unrecognized msg")
	}
	if cmd != nil {
		t.Fatalf("cmd returned from Update method when msg not recognized")
	}
}
