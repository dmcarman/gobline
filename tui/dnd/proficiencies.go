package dnd

import (
	"github.com/dmcarman/gobline/tui/common/views"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Proficiencies struct {
	Armor     tea.Model
	Weapons   tea.Model
	Tools     tea.Model
	Languages tea.Model
}

func NewProficiencies() *Proficiencies {
	armor, _ := views.NewItemlist("Armor", []string{"Light Armor", "Medium Armor", "Shields"})
	weapons, _ := views.NewItemlist("Weapons", []string{"Martial Weapons", "Simple Weapons"})
	tools, _ := views.NewItemlist("Tools", []string{"Navigator's Tools"})
	languages, _ := views.NewItemlist("Languages", []string{"Common"})
	return &Proficiencies{
		Armor:     armor,
		Weapons:   weapons,
		Tools:     tools,
		Languages: languages,
	}

}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (p Proficiencies) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (p Proficiencies) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	default:
		return p, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (p Proficiencies) View() string {
	block := lipgloss.JoinVertical(lipgloss.Left, p.Armor.View(), p.Weapons.View(), p.Tools.View(), p.Languages.View())
	return lipgloss.PlaceHorizontal(23, lipgloss.Left, block)
}
