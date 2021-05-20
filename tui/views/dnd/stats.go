package dnd

import (
	"gobline/tui/constants"
	"gobline/tui/views"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Stats struct {
	strength     tea.Model
	dexterity    tea.Model
	constitution tea.Model
	intelligence tea.Model
	wisdom       tea.Model
	charisma     tea.Model
}

type StatsUpdateMsg struct {
	Label string
	Value int
	Max   int
}

func NewStats() *Stats {
	// strength, _ := views.NewStatPercentbar(0.0, 10, "#ff6246", "#46ffb5", constants.Strength)
	// dexterity, _ := views.NewStatPercentbar(0.0, 10, "#ff1d37", "#37ff1d", constants.Dexterity)
	// constitution, _ := views.NewStatPercentbar(0.0, 10, "#d64fff", "#ffea4f", constants.Constitution)
	// intelligence, _ := views.NewStatPercentbar(0.0, 10, "#ff6726", "#26d4ff", constants.Intelligence)
	// wisdom, _ := views.NewStatPercentbar(0.0, 10, "#ff6726", "#26d4ff", constants.Wisdom)
	// charisma, _ := views.NewStatPercentbar(0.0, 10, "#ff6726", "#26d4ff", constants.Charisma)
	strength, _ := views.NewStatPercentbar(0.0, 10, constants.Strength)
	dexterity, _ := views.NewStatPercentbar(0.0, 10, constants.Dexterity)
	constitution, _ := views.NewStatPercentbar(0.0, 10, constants.Constitution)
	intelligence, _ := views.NewStatPercentbar(0.0, 10, constants.Intelligence)
	wisdom, _ := views.NewStatPercentbar(0.0, 10, constants.Wisdom)
	charisma, _ := views.NewStatPercentbar(0.0, 10, constants.Charisma)

	return &Stats{
		strength:     strength,
		dexterity:    dexterity,
		constitution: constitution,
		intelligence: intelligence,
		wisdom:       wisdom,
		charisma:     charisma,
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s Stats) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s Stats) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StatsUpdateMsg:
		switch msg.Label {
		case constants.Strength:
			v, _ := s.strength.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.strength = v
			return s, nil
		case constants.Dexterity:
			v, _ := s.dexterity.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.dexterity = v
			return s, nil
		case constants.Constitution:
			v, _ := s.constitution.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.constitution = v
			return s, nil
		case constants.Intelligence:
			v, _ := s.intelligence.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.intelligence = v
			return s, nil
		case constants.Wisdom:
			v, _ := s.wisdom.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.wisdom = v
			return s, nil
		case constants.Charisma:
			v, _ := s.charisma.Update(views.StatUpdateMsg(handleConversion(msg)))
			s.charisma = v
			return s, nil
		default:
			return s, nil
		}
	default:
		return s, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s Stats) View() string {
	return lipgloss.JoinVertical(lipgloss.Right,
		s.strength.View(),
		s.dexterity.View(),
		s.constitution.View(),
		s.intelligence.View(),
		s.wisdom.View(),
		s.charisma.View())
}

func handleConversion(msg StatsUpdateMsg) float64 {
	if msg.Value < 0 {
		return 0.0
	} else if msg.Value > msg.Max {
		return 1.0
	}
	return float64(msg.Value) / float64(msg.Max)
}
