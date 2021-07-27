package views

import (
	"github.com/dmcarman/gobline/tui/common/themes"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Itemlist struct {
	Name             string
	Items            []string
	nameStyle        lipgloss.Style
	descriptionStyle lipgloss.Style
}
type ItemlistOption func(*Itemlist) error

func NewItemlist(name string, items []string, opts ...ItemlistOption) (*Itemlist, error) {
	total := &Itemlist{
		Name:             name,
		Items:            items,
		nameStyle:        themes.ItemlistNameTheme,
		descriptionStyle: themes.ItemlistDescriptionTheme,
	}
	for _, opt := range opts {
		if err := opt(total); err != nil {
			return nil, err
		}
	}

	return total, nil
}

func WithItemlistNameStyle(style lipgloss.Style) ItemlistOption {
	return func(s *Itemlist) error {
		s.nameStyle = style
		return nil
	}
}

func WithItemlistDescriptionStyle(style lipgloss.Style) ItemlistOption {
	return func(s *Itemlist) error {
		s.descriptionStyle = style
		return nil
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (i Itemlist) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (i Itemlist) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	default:
		return i, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (i Itemlist) View() string {
	var items string
	for index, item := range i.Items {
		if index == 0 {
			items = lipgloss.JoinVertical(lipgloss.Left, i.descriptionStyle.MaxWidth(23).Render(item))
		} else {
			items = lipgloss.JoinVertical(lipgloss.Left, items, i.descriptionStyle.MaxWidth(23).Render(item))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, i.nameStyle.Render(i.Name), items)
}
