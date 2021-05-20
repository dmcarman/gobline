package tui

import (
	"fmt"
	"gobline/tui/constants"
	"gobline/tui/themes"
	"gobline/tui/views"
	"gobline/tui/views/dnd"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tui struct {
	header     tea.Model
	options    tea.Model
	notes      tea.Model
	statPanel  tea.Model
	skillPanel tea.Model
	theme      themes.Theme
}

func NewTUI() (*tui, error) {

	dndStats := dnd.NewStats()
	statPanel := views.NewPanel("Stats", dndStats)
	dndSkills := dnd.NewSkills()
	skillPanel := views.NewPanel("Skills", dndSkills)
	notes, err := views.NewNotes()
	if err != nil {
		fmt.Println("Could not initialize notes:", err)
		os.Exit(1)
	}
	options := []views.ButtonAction{
		{
			Label: "Option 1",
			Action: func() tea.Msg {
				time.Sleep(time.Second * 3)
				return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("6"), Alt: false}
			}},
		{
			Label: "Option 2",
			Action: func() tea.Msg {
				time.Sleep(time.Second * 3)
				return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("7"), Alt: false}
			}},
		{
			Label: "Option 3",
			Action: func() tea.Msg {
				time.Sleep(time.Second * 3)
				return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("8"), Alt: false}
			}},
	}
	return &tui{
		header:     views.NewHeaderView("Welcome to Gobline, the command line character sheet for dnd 5e"),
		options:    views.NewHorizontalOptionsView(0, options),
		notes:      notes,
		statPanel:  statPanel,
		skillPanel: skillPanel,
		theme:      themes.NewDefaultTheme(),
	}, nil
}

func (t tui) Init() tea.Cmd {
	return nil
}

func (t tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		n, cmd := t.notes.Update(msg)
		t.notes = n
		return t, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return t, tea.Quit
		case "1":
			ops, cmd := t.options.Update(views.OptionsSelectMsg(1))
			t.options = ops
			return t, cmd
		case "2":
			ops, cmd := t.options.Update(views.OptionsSelectMsg(2))
			t.options = ops
			return t, cmd
		case "3":
			ops, cmd := t.options.Update(views.OptionsSelectMsg(3))
			t.options = ops
			return t, cmd
		case "4":
			ops, cmd := t.options.Update(views.OptionsSelectMsg(4))
			t.options = ops
			return t, cmd
		case "5":
			ops, cmd := t.options.Update(views.OptionsSelectMsg(5))
			t.options = ops
			return t, cmd
		case " ":
			ops, cmd := t.options.Update(views.OptionsActivateMsg{})
			t.options = ops
			return t, cmd
		case "6":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Strength, Value: 15, Max: 20})
			t.statPanel = s
			return t, cmd
		case "7":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Dexterity, Value: 10, Max: 20})
			t.statPanel = s
			return t, cmd
		case "8":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Constitution, Value: 5, Max: 20})
			t.statPanel = s
			return t, cmd
		case "9":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Intelligence, Value: 0, Max: 20})
			t.statPanel = s
			return t, cmd
		case "0":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Wisdom, Value: 40, Max: 20})
			t.statPanel = s
			return t, cmd
		case "-":
			s, cmd := t.statPanel.Update(dnd.StatsUpdateMsg{Label: constants.Charisma, Value: -2, Max: 20})
			t.statPanel = s
			return t, cmd
		default:
			vp, cmd := t.notes.Update(msg)
			t.notes = vp
			return t, cmd
		}
	default:
		return t, nil
	}
}

func (t tui) View() string {
	header := t.header.View()
	body := lipgloss.JoinHorizontal(lipgloss.Top, t.statPanel.View(), t.skillPanel.View(), t.notes.View())
	full := lipgloss.JoinVertical(lipgloss.Left, header, body, t.options.View())
	return lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).Render(full)
}
