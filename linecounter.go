package t

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var LineCounter_Style = lipgloss.NewStyle().
	Align(lipgloss.Right).
	PaddingRight(1).
	Foreground(lipgloss.Color("#ccc")).
	Background(lipgloss.Color("#111"))

func renderLineCounter(currentLine, maxSize int, selected bool) string {
	var renderedLineCounter string
	cLineString := "~"
	if currentLine >= 0 {
		cLineString = strconv.Itoa(currentLine)
	}

	renderedLineCounter = LineCounter_Style.Width(maxSize + 2).Render(cLineString)
	return renderedLineCounter
}
