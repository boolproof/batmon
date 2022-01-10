package main

import (
	"time"

	"github.com/batmon/ui"
	termbox "github.com/nsf/termbox-go"
)

const ver = "1.0"
const timeFmt = "15:04:05"
const timeInterval = 1

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	m := NewModel()
	m.Redraw()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)

loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyCtrlQ:
					break loop
				}

			}
		case <-ticker.C:
			data := make([]ui.DataItem, 0)
			data = append(data, ui.DataItem{
				Label: "Time",
				Value: time.Now().Format(timeFmt),
			})

			m.dataPanel.SetData(data)
			m.Redraw()
		}
	}
}

func NewModel() Model {
	vw, vh := termbox.Size()

	data := make([]ui.DataItem, 0)
	data = append(data, ui.DataItem{
		Label: "Time",
		Value: time.Now().Format(timeFmt),
	})

	m := Model{
		vw:        vw,
		vh:        vh,
		statusBar: ui.NewStatusBar(vw, " BatMon "+ver+"   github.com/boolproof/batmon", termbox.ColorBlack, termbox.ColorCyan),
		dataPanel: ui.NewDataPanel(2, 2, data, termbox.ColorDefault, termbox.ColorDefault),
	}

	return m
}

type Model struct {
	vw        int
	vh        int
	statusBar ui.StatusBar
	dataPanel ui.DataPanel
}

func (m *Model) Redraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	m.statusBar.Redraw()
	m.dataPanel.Redraw()
	termbox.Flush()
}
