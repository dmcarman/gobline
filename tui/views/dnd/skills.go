package dnd

import (
	"gobline/tui/constants"
	"gobline/tui/views"

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
	pad := constants.MaxSkillStringLength
	acrobatics, _ := views.NewSkill(constants.Acrobatics, 2, pad-len(constants.Acrobatics))
	animalhandling, _ := views.NewSkill(constants.AnimalHandling, -10, pad-len(constants.AnimalHandling))
	arcana, _ := views.NewSkill(constants.Arcana, 8, pad-len(constants.Arcana))
	athletics, _ := views.NewSkill(constants.Athletics, -1, pad-len(constants.Athletics))
	deception, _ := views.NewSkill(constants.Deception, 20, pad-len(constants.Deception))
	history, _ := views.NewSkill(constants.History, 15, pad-len(constants.History))
	insight, _ := views.NewSkill(constants.Insight, 11, pad-len(constants.Insight))
	intimidation, _ := views.NewSkill(constants.Intimidation, 6, pad-len(constants.Intimidation))
	investigation, _ := views.NewSkill(constants.Investigation, 7, pad-len(constants.Investigation))
	medicine, _ := views.NewSkill(constants.Medicine, 1, pad-len(constants.Medicine))
	nature, _ := views.NewSkill(constants.Nature, 0, pad-len(constants.Nature))
	perception, _ := views.NewSkill(constants.Perception, 0, pad-len(constants.Perception))
	performance, _ := views.NewSkill(constants.Performance, -5, pad-len(constants.Performance))
	persuasion, _ := views.NewSkill(constants.Persuasion, 20, pad-len(constants.Persuasion))
	religion, _ := views.NewSkill(constants.Religion, 2, pad-len(constants.Religion))
	sleightofhand, _ := views.NewSkill(constants.SleightOfHand, 11, pad-len(constants.SleightOfHand))
	stealth, _ := views.NewSkill(constants.Stealth, 2, pad-len(constants.Stealth))
	survival, _ := views.NewSkill(constants.Survival, 5, pad-len(constants.Survival))
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
