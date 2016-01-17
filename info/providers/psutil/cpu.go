package psutil

import (
	"errors"

	"github.com/nproc/rpi-board-info/info"
	"github.com/shirou/gopsutil/cpu"
)

var (
	ErrNoCPU = errors.New("can't load CPU info")
)

func NewCPUInformationProvider() *CPUInformationProvider {
	return &CPUInformationProvider{}
}

type CPUInformationProvider struct {
}

func (p *CPUInformationProvider) CPU() (info.CPU, error) {
	stat, err := cpu.CPUInfo()
	if err != nil {
		return nil, err
	}
	if len(stat) == 0 {
		return nil, ErrNoCPU
	}
	return info.NewCPU(
		stat[0].ModelName,
		len(stat),
	), nil
}
