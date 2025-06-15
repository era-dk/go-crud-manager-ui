package manager

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	Create = &CreateLayer{
		Label: "Create record",
	}
)

type CreateLayer struct {
	Label string
}

/* LayerInterface */
func (l *CreateLayer) Load() {}

/* LayerInterface */
func (l *CreateLayer) Watch(msg tea.Msg) tea.Cmd {
	if len(Form.Fields) < 1 {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			switch msg.String() {
			case "c":
				SetCurrentPage(&CreatePage)
			}
		}
	}
	return nil
}

/* LayerInterface */
func (l CreateLayer) RenderBody() string {
	return ""
}

/* LayerInterface */
func (l CreateLayer) Help() []HelpCmd {
	if len(Form.Fields) > 0 {
		return []HelpCmd{
			{Label: l.Label, Cmd: "c"},
		}
	}
	return []HelpCmd{}
}