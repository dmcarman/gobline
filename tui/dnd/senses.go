package dnd

import (
	"gobline/tui/common/views"
	"gobline/tui/dnd/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Senses struct {
	Acrobatics     tea.Model
	AnimalHandling tea.Model
	Arcana         tea.Model
	Athletics      tea.Model
	Deception      tea.Model
	History        tea.Model
	Insight        tea.Model
	Intimidation   tea.Model
	Investigation  tea.Model
	Medicine       tea.Model
	Nature         tea.Model
	Perception     tea.Model
	Performance    tea.Model
	Persuasion     tea.Model
	Religion       tea.Model
	SleightOfHand  tea.Model
	Stealth        tea.Model
	Survival       tea.Model
}

func NewSenses() *Senses {
	senseLabelWidthOpt := views.WithMaxLabelWidth(constants.MaxSkillStringLength + 3)
	senseAmountLabelOpt := views.WithMaxAmountWidth(3)
	perception, _ := views.NewTotal(constants.Perception, 15, senseLabelWidthOpt, senseAmountLabelOpt, views.WithColor(constants.DexterityColor))
	investigation, _ := views.NewTotal(constants.Investigation, 12, senseLabelWidthOpt, senseAmountLabelOpt, views.WithColor(constants.WisdomColor))
	insight, _ := views.NewTotal(constants.Insight, 8, senseLabelWidthOpt, senseAmountLabelOpt, views.WithColor(constants.WisdomColor))

	return &Senses{
		Perception:    perception,
		Investigation: investigation,
		Insight:       insight,
	}

}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s Senses) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s Senses) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	default:
		return s, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s Senses) View() string {
	return lipgloss.JoinVertical(lipgloss.Right,
		s.Perception.View(),
		s.Investigation.View(),
		s.Insight.View())
}
