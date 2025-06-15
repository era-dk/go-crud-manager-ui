package manager

import (
	"bytes"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case BreakMsg:
		return m, nil
	}

	for _, layer := range CurrentPage.Layers {
		if cmd := layer.Watch(msg); cmd != nil {
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Batch(
				tea.ExitAltScreen,
				tea.Quit,
			)
		}
	}

	return m, nil
}

func (m model) View() string {
	var buffer bytes.Buffer

	if CurrentPage.Title != "" {
		buffer.WriteString(TitleStyle.Render(CurrentPage.Title))
		buffer.WriteString("\n")
	}

	buffer.WriteString("\n")
	for _, layer := range CurrentPage.Layers {
		body := strings.TrimRight(layer.RenderBody(), "\n")
		if body != "" {
			buffer.WriteString(body)
			buffer.WriteString("\n")
		}
	}

	var guide []HelpCmd
	for _, layer := range CurrentPage.Layers {
		for _, help := range layer.Help() {
			guide = append(guide, help)
		}
	}
	guide = append(guide, HelpCmd{Label: "Exit", Cmd: "esc / ctrl+c"})

	buffer.WriteString("\n")
	for _, help := range guide {
		buffer.WriteString(help.Render())
		buffer.WriteString("\n")
	}
	buffer.WriteString("\n")
	
	return buffer.String()
}