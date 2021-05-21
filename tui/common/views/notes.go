package views

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Notes struct {
	vp viewport.Model
}

const content = `
# Some markdown content should go here
---
1. a list
1. of items
1. or whatever

## sub heading
Lorem ipsum dolor sit amet, consectetur adipiscing elit. In at diam non metus faucibus iaculis in a orci. Etiam viverra iaculis ultricies. Donec sollicitudin, libero quis congue aliquet, diam odio convallis lacus, sit amet pellentesque erat arcu quis eros. Nam vitae lectus eget nibh blandit finibus at id tortor. Pellentesque sem eros, laoreet sit amet cursus a, convallis vestibulum dui. Ut non mi arcu. Donec vel augue tellus. Praesent facilisis dui ornare mauris luctus elementum. Aliquam fermentum leo eu metus vestibulum, eu tristique nibh maximus. Ut ac quam non ante porttitor tincidunt sed sed purus.

### sub sub heading
Praesent vel egestas dui, luctus auctor odio. Duis ligula dolor, luctus vel ante sed, aliquam finibus lacus. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed id volutpat nisl, sagittis rhoncus sapien. Donec sed dui vitae metus elementum sagittis vitae a tellus. Phasellus egestas id libero quis sagittis. Maecenas nulla lacus, bibendum non risus a, bibendum imperdiet urna. Sed pharetra, dolor at molestie suscipit, lectus ex laoreet nulla, auctor suscipit erat leo nec mi. Phasellus eget pharetra tellus. Nullam lectus tortor, faucibus sit amet euismod id, hendrerit quis arcu. Ut euismod condimentum augue vitae interdum.  

`

func NewNotes() (*Notes, error) {
	vp := viewport.Model{Width: 20, Height: 24}

	renderer, err := glamour.NewTermRenderer(glamour.WithStylePath("dark"), glamour.WithColorProfile(lipgloss.ColorProfile()))
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}
	vp.SetContent(str)
	return &Notes{vp}, nil
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (n *Notes) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (n *Notes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// n.vp.Width = msgT.Width
		return n, nil
	case tea.KeyMsg:
		switch msg.String() {
		default:
			vp, cmd := n.vp.Update(msg)
			n.vp = vp
			return n, cmd
		}
	default:
		return n, nil
	}
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (n *Notes) View() string {
	return n.vp.View()
}
