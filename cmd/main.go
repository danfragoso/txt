package main

import (
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	t "github.com/danfragoso/txt"
)

type EditorModel struct {
	Content string
	Editor  *t.Editor
}

func (m *EditorModel) Init() tea.Cmd {
	return nil
}

func (m *EditorModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.Editor.SetSize(msg.Width, msg.Height)

	case tea.MouseMsg:
		mouseEvt := tea.MouseEvent(msg)
		switch mouseEvt.Type {
		case tea.MouseWheelDown:
			//_, y := m.Editor.SelectedBuffer().ScrollOffset()
			//m.Editor.SelectedBuffer().SetYScrollOffset(y - 1)
			m.Editor.SelectedBuffer().MoveCursor("down")

		case tea.MouseWheelUp:
			//_, y := m.Editor.SelectedBuffer().ScrollOffset()
			//m.Editor.SelectedBuffer().SetYScrollOffset(y + 1)
			m.Editor.SelectedBuffer().MoveCursor("up")

		case tea.MouseLeft:
			t.Log(mouseEvt)
			lineCounterWidth := len(strconv.Itoa(m.Editor.SelectedBuffer().LineCounter())) + 2
			_, y := m.Editor.SelectedBuffer().ScrollOffset()
			m.Editor.SelectedBuffer().SetCursorX(mouseEvt.X - lineCounterWidth - 1)
			m.Editor.SelectedBuffer().SetCursorY(mouseEvt.Y + y)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "ctrl+x":
			m.Editor.Save()
			return m, tea.Quit

		case "up", "down", "left", "right":
			m.Editor.SelectedBuffer().MoveCursor(msg.String())

		case "esc":
			m.Editor.ToogleMode()

		case "ctrl+s":
			m.Editor.Save()

		default:
			if m.Editor.Mode() == t.EDIT_MODE && msg.Type == tea.KeyRunes && !msg.Alt {
				m.Editor.SelectedBuffer().AddRunes(msg.Runes)
			} else {
				m.Editor.HandleKey(msg.String())
			}

			m.Content = msg.String()
		}
	}

	return m, nil
}

func (m *EditorModel) View() string {
	var editorView string

	editorView += t.Viewport(m.Editor)
	editorView += t.StatusBar(m.Editor)

	return editorView
}

func main() {
	filename := os.Args[1]
	model := &EditorModel{
		Editor: t.CreateEditor(filename),
	}

	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
