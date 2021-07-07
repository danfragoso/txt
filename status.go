package t

import "github.com/charmbracelet/lipgloss"

var StatusTag_Style = lipgloss.NewStyle().
	Bold(true).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Width(10)

var StatusBar_Style = lipgloss.NewStyle().
	Align(lipgloss.Right).
	Foreground(lipgloss.Color("#ccc")).
	Background(lipgloss.Color("#111")).
	Height(1).
	PaddingRight(2)

func StatusBar(editor *Editor) string {
	var statusBar string
	w, _ := editor.Size()

	statusBar += StatusTag_Style.Render(editor.mode.String())
	statusBar += StatusBar_Style.
		Width(w - StatusTag_Style.GetWidth()).
		Render(editor.StatusString())

	return statusBar
}
