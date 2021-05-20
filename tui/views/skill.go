package views

import (
	"fmt"
	"gobline/tui/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Skill struct {
	Name string
	Mod  int
	pad  int
}
type SkillModUpdateMsg int

func NewSkill(name string, mod int, pad int) (*Skill, error) {
	return &Skill{
		Name: name,
		Mod:  mod,
		pad:  pad,
	}, nil
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s Skill) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s Skill) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SkillModUpdateMsg:
		s.Mod = int(msg)
		return s, nil
	default:
		return s, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s Skill) View() string {
	// namepad := lipgloss.NewStyle().PaddingRight(s.pad).Render(s.Name)
	modstring := fmt.Sprintf("%v", s.Mod)
	namepad := lipgloss.PlaceHorizontal(constants.MaxSkillStringLength, lipgloss.Left, s.Name, lipgloss.WithWhitespaceChars("."))
	modpad := lipgloss.PlaceHorizontal(4, lipgloss.Right, modstring, lipgloss.WithWhitespaceChars(" "))
	return lipgloss.JoinHorizontal(lipgloss.Center, namepad, modpad)
}
