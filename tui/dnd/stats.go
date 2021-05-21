package dnd

import (
	"gobline/tui/common/views"
	"gobline/tui/dnd/constants"

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
}

func NewStats() *Stats {
	strength := views.NewStat(0, 20)
	dexterity := views.NewStat(0, 20)
	constitution := views.NewStat(0, 20)
	intelligence := views.NewStat(0, 20)
	wisdom := views.NewStat(0, 20)
	charisma := views.NewStat(0, 20)
	statProgressWidthOpt := views.WithProgressWidth(10)
	strengthBar, _ := views.NewStatPercentbar(strength, constants.Strength, views.WithStatLabelColor(constants.StrengthColor), views.WithProgressSolidFill(constants.StrengthColor), statProgressWidthOpt)
	dexterityBar, _ := views.NewStatPercentbar(dexterity, constants.Dexterity, views.WithStatLabelColor(constants.DexterityColor), views.WithProgressSolidFill(constants.DexterityColor), statProgressWidthOpt)
	constitutionBar, _ := views.NewStatPercentbar(constitution, constants.Constitution, views.WithStatLabelColor(constants.ConstitutionColor), views.WithProgressSolidFill(constants.ConstitutionColor), statProgressWidthOpt)
	intelligenceBar, _ := views.NewStatPercentbar(intelligence, constants.Intelligence, views.WithStatLabelColor(constants.IntelligenceColor), views.WithProgressSolidFill(constants.IntelligenceColor), statProgressWidthOpt)
	wisdomBar, _ := views.NewStatPercentbar(wisdom, constants.Wisdom, views.WithStatLabelColor(constants.WisdomColor), views.WithProgressSolidFill(constants.WisdomColor), statProgressWidthOpt)
	charismaBar, _ := views.NewStatPercentbar(charisma, constants.Charisma, views.WithStatLabelColor(constants.CharismaColor), views.WithProgressSolidFill(constants.CharismaColor), statProgressWidthOpt)

	return &Stats{
		strength:     strengthBar,
		dexterity:    dexterityBar,
		constitution: constitutionBar,
		intelligence: intelligenceBar,
		wisdom:       wisdomBar,
		charisma:     charismaBar,
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
			v, _ := s.strength.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
			s.strength = v
			return s, nil
		case constants.Dexterity:
			v, _ := s.dexterity.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
			s.dexterity = v
			return s, nil
		case constants.Constitution:
			v, _ := s.constitution.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
			s.constitution = v
			return s, nil
		case constants.Intelligence:
			v, _ := s.intelligence.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
			s.intelligence = v
			return s, nil
		case constants.Wisdom:
			v, _ := s.wisdom.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
			s.wisdom = v
			return s, nil
		case constants.Charisma:
			v, _ := s.charisma.Update(views.StatPercentBarUpdateMsg(int(msg.Value)))
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
