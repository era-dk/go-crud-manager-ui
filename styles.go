package manager

import (
	"github.com/charmbracelet/lipgloss"
)

var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("9"))

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Underline(true).
	Foreground(lipgloss.Color("5"))

var TableBaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var PagerActiveStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f9ff4f"))
var PagerDisabledStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#797979"))
	
var GuideLabelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#888"))
var GuideCmdStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#d6d6d6"))

var FieldStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))
var FocusedFieldStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("12"))
var SubmitStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground(lipgloss.Color("#9d3b0f"))
var FocusedSubmitStyle = lipgloss.NewStyle().
	Bold(true).
	Underline(true).
	Foreground(lipgloss.Color("202"))