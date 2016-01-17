package psutil

import (
	"github.com/nproc/rpi-board-info/info"
	"github.com/shirou/gopsutil/load"
)

func NewLoadAverageInformationProvider() *LoadAverageInformationProvider {
	return &LoadAverageInformationProvider{}
}

type LoadAverageInformationProvider struct {
}

func (p *LoadAverageInformationProvider) LoadAverage() (info.LoadAverage, error) {
	stat, err := load.LoadAvg()
	if err != nil {
		return nil, err
	}
	return info.NewLoadAverage(
		stat.Load1,
		stat.Load5,
		stat.Load15,
	), nil
}
