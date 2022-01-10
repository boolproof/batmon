package ui

import (
	termbox "github.com/nsf/termbox-go"
)

func NewStatusBar(w int, text string, fgC, bgC termbox.Attribute) StatusBar {
	e := StatusBar{
		w:    w,
		text: text,
		fgC:  fgC,
		bgC:  bgC,
	}

	return e
}

type StatusBar struct {
	w    int
	text string
	fgC  termbox.Attribute
	bgC  termbox.Attribute
}

func (b *StatusBar) Redraw() {
	i := 0
	for _, char := range b.text {
		termbox.SetCell(i, 0, rune(char), b.fgC, b.bgC)
		i++
	}

	for i < b.w {
		termbox.SetCell(i, 0, rune(' '), b.fgC, b.bgC)
		i++
	}
}
