package manager

import tea "github.com/charmbracelet/bubbletea"

type FormFieldInterface interface {
	Load()
	Focus()
	Blur()
	Key() string
	GetValue() FormValue
	SetValue(v FormValue)
	Validate() error
	Watch(tea.Msg) tea.Cmd
	Render() string
}