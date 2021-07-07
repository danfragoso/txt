package t

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Line struct {
	content string
}

func (l *Line) Lenght() int {
	return len(l.content)
}

func (l *Line) Replace(c string) {
	l.content = c
}

func (l *Line) Content() string {
	return l.content
}

func (l *Line) Append(c string) string {
	l.content += c
	return l.content
}

func (l *Line) Prepend(c string) string {
	l.content = c + l.content
	return l.content
}

func (l *Line) AddAtIndex(c string, idx int) {
	if idx == l.Lenght() {
		l.Append(c)
		return
	}

	l.content = l.content[0:idx] + c + l.content[idx:]
}

func (l *Line) RemoveAtIndex(idx int) {
	l.content = l.content[:idx-1] + l.content[idx:]
}

func (l *Line) SplitAtIndex(idx int) (string, string) {
	if idx == l.Lenght() {
		return l.content, ""
	}

	return l.content[0:idx], l.content[idx:]
}

func readLinesFromPath(path string) []*Line {
	fileStr, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var lines []*Line
	strLines := strings.Split(string(fileStr), "\n")
	for _, line := range strLines {
		lines = append(lines, &Line{line})
	}

	return lines
}
