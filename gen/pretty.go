package gen

import (
	"github.com/charmbracelet/lipgloss"
)

var Style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#D8BFD8"))

var Bold = lipgloss.NewStyle().
	Bold(true)

var Green = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#93C572"))

var StylePlainBorder = lipgloss.NewStyle().
	Border(PlainBorder).
	Padding(1)

var StyleCuteBorder = lipgloss.NewStyle().
	Border(CuteBorder).
	Padding(2).
	Width(76)

var BasicPadding = lipgloss.NewStyle().
	Padding(1)

var CuteBorder = lipgloss.Border{
	Top:         "._.:*:",
	Bottom:      "._.:*:",
	Left:        "|*",
	Right:       "|*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

var PlainBorder = lipgloss.Border{
	Top:         "-",
	Bottom:      "-",
	Left:        "|",
	Right:       "|",
	TopLeft:     ",",
	TopRight:    ",",
	BottomLeft:  "*",
	BottomRight: "*",
}
