package dnd

import (
	"gobline/tui/common/views"
	"gobline/tui/dnd/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Skills struct {
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

func NewSkills() *Skills {
	skillLabelWidthOpt := views.WithMaxLabelWidth(constants.MaxSkillStringLength)
	skillAmountLabelOpt := views.WithMaxAmountWidth(6)
	skillFormatStringOpt := views.WithTotalFormatString("[%+3d]")
	acrobatics, _ := views.NewTotal(constants.Acrobatics, 2, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.DexterityColor), skillFormatStringOpt)
	animalhandling, _ := views.NewTotal(constants.AnimalHandling, -10, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.WisdomColor), skillFormatStringOpt)
	arcana, _ := views.NewTotal(constants.Arcana, 8, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.IntelligenceColor), skillFormatStringOpt)
	athletics, _ := views.NewTotal(constants.Athletics, -1, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.StrengthColor), skillFormatStringOpt)
	deception, _ := views.NewTotal(constants.Deception, 20, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.CharismaColor), skillFormatStringOpt)
	history, _ := views.NewTotal(constants.History, 15, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.IntelligenceColor), skillFormatStringOpt)
	insight, _ := views.NewTotal(constants.Insight, 11, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.WisdomColor), skillFormatStringOpt)
	intimidation, _ := views.NewTotal(constants.Intimidation, 6, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.CharismaColor), skillFormatStringOpt)
	investigation, _ := views.NewTotal(constants.Investigation, 7, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.IntelligenceColor), skillFormatStringOpt)
	medicine, _ := views.NewTotal(constants.Medicine, 1, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.WisdomColor), skillFormatStringOpt)
	nature, _ := views.NewTotal(constants.Nature, 0, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.IntelligenceColor), skillFormatStringOpt)
	perception, _ := views.NewTotal(constants.Perception, 0, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.WisdomColor), skillFormatStringOpt)
	performance, _ := views.NewTotal(constants.Performance, -5, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.CharismaColor), skillFormatStringOpt)
	persuasion, _ := views.NewTotal(constants.Persuasion, 20, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.CharismaColor), skillFormatStringOpt)
	religion, _ := views.NewTotal(constants.Religion, 2, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.IntelligenceColor), skillFormatStringOpt)
	sleightofhand, _ := views.NewTotal(constants.SleightOfHand, 11, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.DexterityColor), skillFormatStringOpt)
	stealth, _ := views.NewTotal(constants.Stealth, 2, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.DexterityColor), skillFormatStringOpt)
	survival, _ := views.NewTotal(constants.Survival, 5, skillLabelWidthOpt, skillAmountLabelOpt, views.WithColor(constants.WisdomColor), skillFormatStringOpt)
	return &Skills{
		Acrobatics:     acrobatics,
		AnimalHandling: animalhandling,
		Arcana:         arcana,
		Athletics:      athletics,
		Deception:      deception,
		History:        history,
		Insight:        insight,
		Intimidation:   intimidation,
		Investigation:  investigation,
		Medicine:       medicine,
		Nature:         nature,
		Perception:     perception,
		Performance:    performance,
		Persuasion:     persuasion,
		Religion:       religion,
		SleightOfHand:  sleightofhand,
		Stealth:        stealth,
		Survival:       survival,
	}

}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s Skills) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s Skills) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	default:
		return s, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (s Skills) View() string {
	return lipgloss.JoinVertical(lipgloss.Right,
		s.Acrobatics.View(),
		s.AnimalHandling.View(),
		s.Arcana.View(),
		s.Athletics.View(),
		s.Deception.View(),
		s.History.View(),
		s.Insight.View(),
		s.Intimidation.View(),
		s.Investigation.View(),
		s.Medicine.View(),
		s.Nature.View(),
		s.Perception.View(),
		s.Performance.View(),
		s.Persuasion.View(),
		s.Religion.View(),
		s.SleightOfHand.View(),
		s.Stealth.View(),
		s.Survival.View())
}
