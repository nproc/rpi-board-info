package native

import (
	"bytes"
	"io/ioutil"
	"strconv"

	"github.com/nproc/rpi-board-info/info"
)

func NewTemperatureInformationProvider() *TemperatureInformationProvider {
	return &TemperatureInformationProvider{}
}

type TemperatureInformationProvider struct {
}

func (p *TemperatureInformationProvider) CPUTemperature() (info.Temperature, error) {
	stat, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return nil, err
	}
	stat = bytes.TrimSpace(stat)
	value, err := strconv.Atoi(string(stat))
	if err != nil {
		return nil, err
	}
	return info.NewTemperature(value), nil
}
