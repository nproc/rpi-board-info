package psutil

import (
	"github.com/nproc/rpi-board-info/info"
	"github.com/shirou/gopsutil/mem"
)

func NewMemoryInformationProvider() *MemoryInformationProvider {
	return &MemoryInformationProvider{}
}

type MemoryInformationProvider struct {
}

func (p *MemoryInformationProvider) SwapMemory() (info.Memory, error) {
	stat, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	return info.NewMemory(
		stat.Total,
		stat.Used,
		stat.Free,
	), nil
}

func (p *MemoryInformationProvider) VirtualMemory() (info.VirtualMemory, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return info.NewVirtualMemory(
		stat.Total,
		stat.Used,
		stat.Free,
		stat.Cached,
	), nil
}
