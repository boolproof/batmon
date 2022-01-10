package ui

import (
	termbox "github.com/nsf/termbox-go"
)

type DataItem struct {
	Label string
	Value string
}

func NewDataPanel(x, y int, data []DataItem, fgC, bgC termbox.Attribute) DataPanel {
	e := DataPanel{
		x:    x,
		y:    y,
		data: data,
		fgC:  fgC,
		bgC:  bgC,
	}

	return e
}

type DataPanel struct {
	x    int
	y    int
	data []DataItem
	fgC  termbox.Attribute
	bgC  termbox.Attribute
}

func (p *DataPanel) Redraw() {
	for y, d := range p.data {
		i := 0
		for _, char := range d.Label + ": " + d.Value {
			termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
			i++
		}
	}
}

func (p *DataPanel) SetData(data []DataItem) {
	p.data = data
}
