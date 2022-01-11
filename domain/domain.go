package domain

import (
	"time"

	"github.com/distatus/battery"
)

const TimeFmt = "15:04:05"

type BatteryStatus struct {
	Time          time.Time
	State         string
	ChargePctg    float64
	Current       float64
	Full          float64
	Design        float64
	ChargeRate    float64
	Voltage       float64
	DesignVoltage float64
}

func ReadBatteryStatus() BatteryStatus {
	batteries, err := battery.GetAll()
	if err != nil {
		panic("Could not get battery info!")
	}

	mainBat := batteries[0]

	return BatteryStatus{
		Time:          time.Now(),
		State:         mainBat.State.String(),
		ChargePctg:    100 * mainBat.Current / mainBat.Full,
		Current:       mainBat.Current,
		Full:          mainBat.Full,
		Design:        mainBat.Design,
		ChargeRate:    mainBat.ChargeRate,
		Voltage:       mainBat.Voltage,
		DesignVoltage: mainBat.DesignVoltage,
	}
}
