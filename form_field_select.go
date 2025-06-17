package manager

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type FormFieldSelectItem struct {
	ID int
	Label string
}

type FormFieldSelect struct {
	Name string
	Label string
	Items []FormFieldSelectItem
	Required bool
	focused bool
	cursor int
	selected int
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Load() {
	ff.cursor = -1
	ff.selected = -1
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Focus() {
	if ff.selected != -1 {
		ff.cursor = ff.selected
	} else {
		ff.cursor = 0
	}
	ff.focused = true
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Blur() {
	ff.cursor = -1
	ff.focused = false
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Key() string {
	return ff.Name
}

/* FormFieldInterface */
func (ff FormFieldSelect) GetValue() any {
	for i, item := range ff.Items {
		if i == ff.selected {
			return item.ID
		} 
	}
	return 0
}

/* FormFieldInterface */
func (ff *FormFieldSelect) SetValue(v any) {
	vv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
	for i, item := range ff.Items {
		if item.ID == vv {
			ff.selected = i
			break
		} 
	}
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Validate() error {
	id := ff.GetValue().(int)
	if ff.Required && id == 0 {
		return errors.New("field is required")
	}
	return nil
}

/* FormFieldInterface */
func (ff *FormFieldSelect) Watch(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if ff.cursor > 0 {
				ff.cursor--
			} else {
				return nil
			}
			return BreakCmd
		case tea.KeyDown:
			if ff.cursor < len(ff.Items) - 1 {
				ff.cursor++
			} else {
				ff.cursor = -1
				return nil
			}
			return BreakCmd
		case tea.KeySpace:
			ff.selected = ff.cursor
			return BreakCmd
		case tea.KeyEnter:
			ff.selected = ff.cursor
		}
	}
	
	return nil
}

/* FormFieldInterface */
func (ff FormFieldSelect) Render() string {
	var buffer bytes.Buffer

	if ff.focused {
		buffer.WriteString(fmt.Sprintf("%s:\n", FocusedFieldStyle.Render(ff.Label)))
	} else {
		buffer.WriteString(fmt.Sprintf("%s:\n", FieldStyle.Render(ff.Label)))
	}
	
	for i, item := range ff.Items {
		cursor := " "
		if ff.cursor == i {
            cursor = ">"
        }

		selected := " "
		if ff.selected == i {
			selected = "•"
		}

		buffer.WriteString(fmt.Sprintf(" %s (%s) %s\n", cursor, selected, item.Label))
	}

	return buffer.String()
}