package manager

import (
	"bytes"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Edit = &EditLayer{
		Label: "Edit record",
	}
)

type EditLayer struct {
	Label string
	Handle func (id int) (Record, error)
	err error
}

/* LayerInterface */
func (l *EditLayer) Load() {}

/* LayerInterface */
func (l *EditLayer) Watch(msg tea.Msg) tea.Cmd {
	if l.Handle == nil || !Grid.HasRows() || len(Form.Fields) < 1 {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			switch msg.String() {
			case "e":
				var record Record
				record, l.err = l.Handle(Grid.FocusedId())
				SetCurrentPage(&EditPage)

				Form.RecordID = record.ID
				for _, field := range Form.Fields {
					field.SetValue(record.Get(field.Key()))
				}
			}
		}
	}
	return nil
}

/* LayerInterface */
func (l EditLayer) RenderBody() string {
	var buffer bytes.Buffer
	if l.err != nil {
		buffer.WriteString(ErrorStyle.Render(l.err.Error()))
		buffer.WriteString("\n")
		l.err = nil
	}
	return buffer.String()
}

/* LayerInterface */
func (l EditLayer) Help() []HelpCmd {
	if l.Handle == nil || !Grid.HasRows() || len(Form.Fields) < 1 {
		return []HelpCmd{}
	}
	
	return []HelpCmd{
		{Label: l.Label, Cmd: "e"},
	}
}