package manager

import tea "github.com/charmbracelet/bubbletea"

type FormFieldInterface interface {
	Load()
	Reset()
	Focus()
	Blur()
	Key() string
	GetValue() any
	SetValue(v any)
	Validate() error
	Watch(tea.Msg) tea.Cmd
	Render() string
}