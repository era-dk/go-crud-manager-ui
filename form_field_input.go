package manager

import (
	"bytes"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FormFieldInput struct {
	Name string
	Label string
	ValidateFn func (value string) error
	model textinput.Model
}

/* FormFieldInterface */
func (ff *FormFieldInput) Load() {
	ff.model = textinput.New()
}

/* FormFieldInterface */
func (ff *FormFieldInput) Focus() {
	ff.model.Focus()
}

/* FormFieldInterface */
func (ff *FormFieldInput) Blur() {
	ff.model.Blur()
}

/* FormFieldInterface */
func (ff *FormFieldInput) Key() string {
	return ff.Name
}

/* FormFieldInterface */
func (ff *FormFieldInput) GetValue() any {
	return ff.model.Value()
}

/* FormFieldInterface */
func (ff *FormFieldInput) SetValue(v any) {
	ff.model.SetValue(fmt.Sprintf("%v", v))
}

/* FormFieldInterface */
func (ff *FormFieldInput) Validate() error {
	if ff.ValidateFn != nil {
		return ff.ValidateFn(ff.model.Value())
	}
	return nil
}

/* FormFieldInterface */
func (ff *FormFieldInput) Watch(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	ff.model, cmd = ff.model.Update(msg)
	return cmd
}

/* FormFieldInterface */
func (ff FormFieldInput) Render() string {
	var buffer bytes.Buffer
	
	if ff.model.Focused() {
		buffer.WriteString(fmt.Sprintf("%s %s \n", FocusedFieldStyle.Render(ff.Label), ff.model.View()))
	} else {
		buffer.WriteString(fmt.Sprintf("%s %s \n", FieldStyle.Render(ff.Label), ff.model.View()))
	}

	return buffer.String()
}