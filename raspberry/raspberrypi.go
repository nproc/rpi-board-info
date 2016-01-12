package raspberry

import (
	"sync"

	"github.com/nproc/rpi-board-info/thermal"
)

type Pi struct {
	mutex       *sync.RWMutex
	thermometer thermal.Thermometer
	Temperature thermal.Temperature `json:"temperature"`
}

func NewPi(thermometer thermal.Thermometer) *Pi {
	return &Pi{
		mutex:       &sync.RWMutex{},
		thermometer: thermometer,
	}
}

func (r *Pi) RecordCurrentTemperature() {
	r.Temperature = r.thermometer.Measure()
}
