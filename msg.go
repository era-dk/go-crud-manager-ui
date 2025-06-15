package manager

import tea "github.com/charmbracelet/bubbletea"

type BreakMsg struct{}

func BreakCmd() tea.Msg {
	return BreakMsg{}
}