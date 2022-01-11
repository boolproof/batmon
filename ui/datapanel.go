package ui

import (
	"fmt"

	"github.com/batmon/domain"
	termbox "github.com/nsf/termbox-go"
)

type DataItem struct {
	Label string
	Value string
}

func NewDataPanel(x, y int, data domain.BatteryStatus, fgC, bgC termbox.Attribute) DataPanel {
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
	data domain.BatteryStatus
	fgC  termbox.Attribute
	bgC  termbox.Attribute
}

func (p *DataPanel) Redraw() {
	i := 0
	y := 0
	for _, char := range "Time: " + p.data.Time.Format(domain.TimeFmt) {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	for _, char := range "Status: " + p.data.State {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	chargePctg := fmt.Sprintf("%f", p.data.ChargePctg)
	for _, char := range "Charge pctg: " + chargePctg + " %" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	current := fmt.Sprintf("%f", p.data.Current)
	for _, char := range "Current: " + current + " mWh" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	full := fmt.Sprintf("%f", p.data.Full)
	for _, char := range "Full: " + full + " mWh" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	design := fmt.Sprintf("%f", p.data.Design)
	for _, char := range "Design: " + design + " mWh" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	chargeRate := fmt.Sprintf("%f", p.data.ChargeRate)
	for _, char := range "Charge rate: " + chargeRate + " mW" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	voltage := fmt.Sprintf("%f", p.data.Voltage)
	for _, char := range "Voltage: " + voltage + " V" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}

	i = 0
	y++
	designVoltage := fmt.Sprintf("%f", p.data.DesignVoltage)
	for _, char := range "Design voltage: " + designVoltage + " V" {
		termbox.SetCell(p.x+i, p.y+y, rune(char), p.fgC, p.bgC)
		i++
	}
}

func (p *DataPanel) SetData(data domain.BatteryStatus) {
	p.data = data
}
