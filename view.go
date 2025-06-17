package manager

import (
	"bytes"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	View = &ViewLayer{
		Label: "View record",
	}
)

type ViewLayer struct {
	Label string
	Handle func (id int) (Record, error)
	Record Record
	err error
}

/* LayerInterface */
func (l *ViewLayer) Load() {
	l.Record = NewRecord(0)
	l.err = nil
}

/* LayerInterface */
func (l *ViewLayer) Watch(msg tea.Msg) tea.Cmd {
	if l.Handle == nil || !Grid.HasRows() {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			switch msg.String() {
			case "v":
				var record Record
				record, l.err = l.Handle(Grid.FocusedId())
				SetCurrentPage(&ViewPage)
				View.Record = record
				return BreakCmd
			}
		case tea.KeyEsc:
			if CurrentPage == &ViewPage {
				SetCurrentPage(&IndexPage)
				return BreakCmd
			}
			return nil
		}
	}
	return nil
}

/* LayerInterface */
func (l ViewLayer) RenderBody() string {
	var buffer bytes.Buffer

	if l.Record.HasKeys() {
		for _, key := range l.Record.Keys() {
			buffer.WriteString(fmt.Sprintf("%s %s \n", FieldStyle.Render(key), l.Record.GetString(key)))
		}
	}

	if l.err != nil {
		buffer.WriteString(ErrorStyle.Render(l.err.Error()))
		buffer.WriteString("\n")
		l.err = nil
	}
	return buffer.String()
}

/* LayerInterface */
func (l ViewLayer) Help() []HelpCmd {
	if l.Handle == nil || !Grid.HasRows() || CurrentPage == &ViewPage {
		return []HelpCmd{}
	}
	
	return []HelpCmd{
		{Label: l.Label, Cmd: "v"},
	}
}