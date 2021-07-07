package t

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var Cursor_Style = lipgloss.NewStyle().
	Reverse(true)

func Viewport(editor *Editor) string {
	var viewport string
	linesToShow := editor.height - StatusBar_Style.GetHeight()
	offset := editor.selectedBuffer.yScrollOffset

	lineCounterWidth := len(strconv.Itoa(editor.selectedBuffer.LineCounter()))
	lineWidth := editor.width - lineCounterWidth

	for i := 0; i < linesToShow; i++ {
		lineIdx := i + offset
		lineHasCursor := lineIdx == editor.selectedBuffer.cursorY

		if lineIdx < editor.selectedBuffer.LineCounter() {
			viewport += renderLineCounter(lineIdx+1, lineCounterWidth, false)
			viewport += renderLine(editor.selectedBuffer,
				editor.selectedBuffer.lines[lineIdx].content,
				lineHasCursor,
				editor.selectedBuffer.cursorX,
				lineWidth,
			)
		}

		viewport += "\n"
	}

	return viewport
}

func renderLine(buffer *Buffer, lineContent string, hasCursor bool, cursorPos int, lineWidth int) string {
	var renderedLine string
	var suffixCursor string

	// If cursor is at the end of line render a suffix cursor
	if hasCursor && cursorPos > len(lineContent)-1 {
		cursorPos = len(lineContent)
		buffer.SetCursorX(cursorPos)
		suffixCursor = Cursor_Style.Render(" ")
	}

	if len(lineContent) > lineWidth {
	}

	for charIdx, r := range lineContent {
		sR := string(r)
		if sR == "\t" {
			sR = "  "
		}

		if hasCursor && charIdx == cursorPos {
			sR = Cursor_Style.Render(sR)
		}

		renderedLine += sR
	}

	return renderedLine + suffixCursor
}
