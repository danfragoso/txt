package t

import (
	"fmt"
)

type EditorMode uint8

func (m EditorMode) String() string {
	return [...]string{"EDIT", "CMD"}[m]
}

const (
	EDIT_MODE EditorMode = iota
	CMD_MODE
)

type Editor struct {
	width  int
	height int

	selectedBuffer *Buffer
	buffers        []*Buffer

	mode EditorMode

	IOStatus string
}

func CreateEditor(filename string) *Editor {
	buff := CreateBuffer(filename)
	editor := &Editor{
		selectedBuffer: buff,
		buffers:        []*Buffer{buff},
	}

	buff.editor = editor
	return editor
}

func (e *Editor) Buffers() []*Buffer {
	return e.buffers
}

func (e *Editor) SetSize(width, height int) {
	e.width, e.height = width, height
}

func (e *Editor) Mode() EditorMode {
	return e.mode
}

func (e *Editor) StatusString() string {
	statusString := "NIL"
	if e.selectedBuffer != nil {
		statusString = e.selectedBuffer.path

		statusString += fmt.Sprintf(" l%d c%d", e.SelectedBuffer().cursorY+1, e.SelectedBuffer().cursorX+1)

	}

	return e.IOStatus + statusString
}

func (e *Editor) ToogleMode() {
	if e.mode == EDIT_MODE {
		e.mode = CMD_MODE
	} else {
		e.mode = EDIT_MODE
	}
}

func (e *Editor) HandleKey(key string) {
	e.IOStatus = "*"

	switch key {
	case "enter":
		e.selectedBuffer.AddNewLine()
	case "backspace":
		e.selectedBuffer.RemoveRune()

	case "tab":
		e.selectedBuffer.AddRunes([]rune{'\t'})
	}
}

func (e *Editor) Save() {
	result := e.selectedBuffer.saveToFile()
	e.IOStatus = result
}

func (e *Editor) Size() (int, int) {
	return e.width, e.height
}

func (e *Editor) SelectedBuffer() *Buffer {
	return e.selectedBuffer
}

func (e *Editor) SelectBufferById(id string) error {
	for _, buffer := range e.buffers {
		if buffer.id == id {
			e.selectedBuffer = buffer
			return nil
		}
	}

	return fmt.Errorf("Buffer %s does not exist on this editor...", id)
}
