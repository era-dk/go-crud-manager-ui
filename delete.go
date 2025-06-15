package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	Delete = &DeleteLayer{
		Label: "Delete record",
		Handle: nil,
	}
)

type DeleteLayer struct {
	Label string
	Handle func (id int) error
	confirm bool
	err error
}

/* LayerInterface */
func (l *DeleteLayer) Load() {}

/* LayerInterface */
func (l *DeleteLayer) Watch(msg tea.Msg) tea.Cmd {
	if l.Handle == nil || !Listing.HasRows() {
		return nil
	}

	if l.confirm {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "y":
				if err := l.Handle(Listing.FocusedId()); err != nil {
					l.err = err
				} else {
					Listing.Load()
				}
			}
		}

		l.confirm = false
		return BreakCmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDelete, tea.KeyBackspace:
			l.confirm = true
		}
	}
	return nil
}

/* LayerInterface */
func (l *DeleteLayer) RenderBody() string {
	if l.confirm {
		tiModel := textinput.New()
		tiModel.Focus()

		return fmt.Sprintf(
			"Delete record [ID: %d] [y/N] %s",
			Listing.FocusedId(),
			tiModel.View(),
		)
	}
	if l.err != nil {
		s := fmt.Sprintf("%s\n", ErrorStyle.Render(l.err.Error()))
		l.err = nil
		return s
	}

	return ""
}

/* LayerInterface */
func (l DeleteLayer) Help() []HelpCmd {
	if l.Handle == nil || !Listing.HasRows() {
		return []HelpCmd{}
	}
	
	return []HelpCmd{
		{Label: l.Label, Cmd: "delete"},
	}
}