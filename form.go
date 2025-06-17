package manager

import (
	"bytes"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Form = &FormLayer{
		RecordID: 0,
		Fields: []FormFieldInterface{},
		focusedIdx: -1,
		errsMap: map[string]error{},
	}
)

type FormLayer struct {
	RecordID int
	Fields []FormFieldInterface
	focusedIdx int
	errsMap map[string]error
}

func (l *FormLayer) Validate() bool {
	valid := true
	l.errsMap = map[string]error{}
	for _, field := range l.Fields {
		if err := field.Validate(); err != nil {
			l.errsMap[field.Key()] = err
			valid = false
		}
	}
	return valid
}

func (l *FormLayer) GetField(name string) FormFieldInterface {
	for _, field := range l.Fields {
		if field.Key() == name {
			return field
		}
	}
	return nil
}

/* LayerInterface */
func (l *FormLayer) Load() {
	l.RecordID = 0
	for _, field := range l.Fields {
		field.Load()
	}
	l.focusedIdx = 0
	l.errsMap = map[string]error{}
	l.focusedField().Focus()
}

/* LayerInterface */
func (l *FormLayer) Watch(msg tea.Msg) tea.Cmd {
	if len(l.Fields) < 1 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEsc:
				SetCurrentPage(&IndexPage)
			}
		}
		return BreakCmd
	}

	if cmd := l.focusedField().Watch(msg); cmd != nil {
		return cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			focusedField := l.focusedField()
			delete(l.errsMap, focusedField.Key())
			if err := focusedField.Validate(); err != nil {
				l.errsMap[focusedField.Key()] = err
			} else {
				l.focusNextField()
			}
		case tea.KeyUp:
			l.focusPrevField()
		case tea.KeyDown, tea.KeyTab:
			l.focusNextField()
		case tea.KeyEsc:
			SetCurrentPage(&IndexPage)
			return BreakCmd
		}
	}

	return nil
}

/* LayerInterface */
func (l FormLayer) RenderBody() string {
	if len(l.Fields) < 1 {
		return ""
	}

	var buffer bytes.Buffer
	
	for _, field := range l.Fields {
		buffer.WriteString(field.Render())
		err, ok := l.errsMap[field.Key()]
		if ok && err != nil {
			buffer.WriteString(ErrorStyle.Render(err.Error()))
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

/* LayerInterface */
func (l FormLayer) Help() []HelpCmd {
	return []HelpCmd{}
}

func (l *FormLayer) focusedField() FormFieldInterface {
	return l.Fields[l.focusedIdx]
}

func (l *FormLayer) focusPrevField() {
	newFocusIdx := l.focusedIdx - 1
	if newFocusIdx > -1 {
		l.focusedField().Blur()
		l.focusedIdx = newFocusIdx
		l.focusedField().Focus()
	}
}

func (l *FormLayer) focusNextField() {
	newFocusIdx := l.focusedIdx + 1
	if newFocusIdx < len(l.Fields) {
		l.focusedField().Blur()
		l.focusedIdx = newFocusIdx
		l.focusedField().Focus()
	}
}