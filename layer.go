package manager

import tea "github.com/charmbracelet/bubbletea"

type LayerInterface interface {
	Load()
	Watch(tea.Msg) tea.Cmd
	RenderBody() string
	Help() []HelpCmd
}