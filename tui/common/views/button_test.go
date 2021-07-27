package views_test

import (
	"gobline/tui/common/themes"
	"gobline/tui/common/views"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type testMsg struct{ tea.Msg }

func getTestButton(f tea.Cmd) tea.Model {
	testTheme := themes.NewDefaultTheme()
	testTheme.ButtonNormal = lipgloss.NewStyle()
	testTheme.ButtonHighlight = lipgloss.NewStyle().Bold(true)
	testTheme.ButtonDisabled = lipgloss.NewStyle().Faint(true)
	btn := views.NewButtonView(false, "Test", f).WithStyle(testTheme)
	return btn
}
func TestButtonSelect(t *testing.T) {
	btn := getTestButton(nil)
	if btn.View() != "Test" {
		t.Fatalf("Button normal style not applied")
	}
	btn, _ = btn.Update(views.ButtonSelectMsg(true))
	if btn.View() != "\x1B[1mTest\x1B[0m" {
		t.Fatalf("Button highlight style not applied after select msg")
	}

}
func TestButtonDisable(t *testing.T) {
	btn := getTestButton(nil)
	if btn.View() != "Test" {
		t.Fatalf("Button normal style not applied")
	}
	btn, _ = btn.Update(views.ButtonDisableMsg(true))
	if btn.View() != "\x1B[2mTest\x1B[0m" {
		t.Fatalf("Button disabled style not applied after disable msg")
	}

}
func TestButtonText(t *testing.T) {
	btn := getTestButton(nil)
	if btn.View() != "Test" {
		t.Fatalf("Button normal style not applied")
	}
	btn, _ = btn.Update(views.ButtonTextMsg("Success"))
	if btn.View() != "Success" {
		t.Fatalf("Button text not updated after text msg sent")
	}

}
func TestButtonActivate(t *testing.T) {
	btn := getTestButton(func() tea.Msg {
		t.Log("cmd function run")
		return views.ButtonTextMsg("Success")
	})
	_, cmd := btn.Update(views.ButtonActivateMsg{})
	if cmd == nil {
		t.Fatalf("cmd not returned from Update method")
	}
	cmd()

}
func TestButtonActivateWhileDisabled(t *testing.T) {
	btn := getTestButton(func() tea.Msg {
		t.Log("cmd function ran successfully")
		return views.ButtonTextMsg("Success")
	})
	btn, _ = btn.Update(views.ButtonDisableMsg(true))
	_, cmd := btn.Update(views.ButtonActivateMsg{})
	if cmd != nil {
		t.Fatalf("cmd returned from Update method when button should be disabled")
	}

}

func TestButtonUnrecognizedMsg(t *testing.T) {
	btn := getTestButton(nil)
	_, cmd := btn.Update(testMsg{})
	if cmd != nil {
		t.Fatalf("cmd returned from Update method when msg not recognized")
	}

}

func TestButtonInit(t *testing.T) {
	btn := getTestButton(nil)
	cmd := btn.Init()
	if cmd != nil {
		t.Fatalf("button init method should return nil")
	}
}
