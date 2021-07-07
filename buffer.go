package t

import "io/ioutil"

type Buffer struct {
	editor *Editor
	id     string

	path  string
	lines []*Line

	xScrollOffset int
	yScrollOffset int

	cursorX int
	cursorY int
}

func CreateBuffer(path string) *Buffer {
	buffer := &Buffer{path: path}
	if buffer.path != "" {
		buffer.lines = readLinesFromPath(buffer.path)
	}

	return buffer
}

func (b *Buffer) LineCounter() int {
	return len(b.lines)
}

func (b *Buffer) LineByIndex(i int) *Line {
	return b.lines[i]
}

func (b *Buffer) Id() string {
	return b.id
}

func (b *Buffer) AddRunes(runes []rune) {
	b.editor.IOStatus = "*"
	b.LineByIndex(b.cursorY).AddAtIndex(string(runes), b.cursorX)
	b.MoveCursor("right")
}

func (b *Buffer) saveToFile() string {
	var fileData []byte
	for idx, line := range b.lines {
		nLine := "\n"
		if idx == len(b.lines)-1 {
			nLine = ""
		}

		byteLine := []byte(line.content + nLine)
		fileData = append(fileData, byteLine...)
	}

	err := ioutil.WriteFile(b.path, fileData, 0644)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (b *Buffer) RemoveRune() {
	currentLine := b.LineByIndex(b.cursorY)
	if b.cursorX == 0 && b.cursorY-1 >= 0 {
		previousLine := b.LineByIndex(b.cursorY - 1)
		b.cursorX = len(previousLine.content)

		b.RemoveLine(b.cursorY)
		previousLine.content += currentLine.content
		b.MoveCursor("up")

		return
	}

	currentLine.RemoveAtIndex(b.cursorX)
	b.MoveCursor("left")
}

func (b *Buffer) RemoveLine(idx int) {
	b.lines = append(b.lines[:idx], b.lines[idx+1:]...)
}

func (b *Buffer) AddNewLine() {
	left, right := b.LineByIndex(b.cursorY).SplitAtIndex(b.cursorX)
	b.LineByIndex(b.cursorY).Replace(left)

	b.AddLineAtIndex(&Line{right}, b.cursorY+1)
	b.cursorX = 0
	b.MoveCursor("down")
}

func (b *Buffer) AddLineAtIndex(line *Line, idx int) {
	if idx >= len(b.lines) {
		b.lines = append(b.lines, &Line{})
		return
	}

	oldLine := b.lines[idx]
	head := append(b.lines[:idx], line)
	b.lines = append(head, b.lines[idx:]...)
	b.lines[idx+1] = oldLine
}

func (b *Buffer) MoveCursor(position string) {
	var scrollBuffer = 2
	switch position {
	case "up":
		if b.cursorY > 0 {
			b.cursorY -= 1
			_, y := b.ScrollOffset()
			if b.cursorY == y+scrollBuffer-1 {
				b.SetYScrollOffset(y - 1)
			}
		}

	case "down":
		if b.cursorY < len(b.lines)-1 {
			b.cursorY += 1
			linesToShow := b.editor.height - StatusBar_Style.GetHeight()
			if b.cursorY == (linesToShow+b.yScrollOffset)-scrollBuffer {
				_, y := b.ScrollOffset()
				b.SetYScrollOffset(y + 1)
			}
		}

	case "left":
		if b.cursorX > 0 {
			b.cursorX -= 1
		} else {
			b.MoveCursor("up")
			b.cursorX = len(b.lines[b.cursorY].content)
		}

	case "right":
		if b.cursorX < len(b.lines[b.cursorY].content) {
			b.cursorX += 1
		} else {
			b.cursorX = 0
			b.MoveCursor("down")
		}
	}
}

func (b *Buffer) SetCursorY(pos int) {
	if pos >= 0 {
		b.cursorY = pos
	}
}

func (b *Buffer) SetCursorX(pos int) {
	if pos >= 0 {
		b.cursorX = pos
	}
}

func (b *Buffer) ScrollOffset() (int, int) {
	return b.xScrollOffset, b.yScrollOffset
}

func (b *Buffer) SetXScrollOffset(offset int) {
	if offset >= 0 {
		b.xScrollOffset = offset
	}
}

func (b *Buffer) SetYScrollOffset(offset int) {
	if offset >= 0 && offset < len(b.lines) {
		b.yScrollOffset = offset
	}
}
