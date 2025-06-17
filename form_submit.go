package manager

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type FormSubmit struct {
	Label string
	Handle func (id int) error
	focused bool
	err error
}

/* FormFieldInterface */
func (ff *FormSubmit) Load() {
	ff.err = nil
}

/* FormFieldInterface */
func (ff *FormSubmit) Focus() {
	ff.focused = true
}

/* FormFieldInterface */
func (ff *FormSubmit) Blur() {
	ff.focused = false
}

/* FormFieldInterface */
func (ff FormSubmit) Key() string {
	return "_submit"
}

/* FormFieldInterface */
func (ff FormSubmit) GetValue() FormValue {
	return FormValue{}
}

/* FormFieldInterface */
func (ff FormSubmit) SetValue(v FormValue) {}

/* FormFieldInterface */
func (ff FormSubmit) Validate() error {
	return ff.err
}

/* FormFieldInterface */
func (ff *FormSubmit) Watch(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if Form.Validate() {
				if ff.Handle != nil {
					if err := ff.Handle(Form.RecordID); err != nil {
						ff.err = err
					} else {
						SetCurrentPage(&IndexPage)
					}
				}
			}
		}
	}

	return nil
}

/* FormFieldInterface */
func (ff FormSubmit) Render() string {
	if ff.focused {
		return fmt.Sprintf("%s\n", FocusedSubmitStyle.Render(fmt.Sprintf("[ %s ]", ff.Label)))
	} else {
		return fmt.Sprintf("%s\n", SubmitStyle.Render(fmt.Sprintf("[ %s ]", ff.Label)))
	}
}