package gql

import (
	"time"

	"github.com/nproc/rpi-board-info/info"
)

func NewBoardData(src info.Board) *BoardData {
	return &BoardData{
		Hostname: src.Hostname(),
		OS:       src.OS(),
		Platform: src.Platform(),
		Family:   src.Family(),
		Version:  src.Version(),
		BootTime: src.BootTime(),
	}
}

type BoardData struct {
	Hostname string    `json:"hostname"`
	OS       string    `json:"os"`
	Platform string    `json:"platform"`
	Family   string    `json:"family"`
	Version  string    `json:"version"`
	BootTime time.Time `json:"boottime"`
}

func NewCPUData(src info.CPU) *CPUData {
	return &CPUData{
		Model: src.Model(),
		Cores: src.Cores(),
	}
}

type CPUData struct {
	Model string `json:"model"`
	Cores int    `json:"cores"`
}

func NewDiskDataList(src []info.Disk) []*DiskData {
	rs := make([]*DiskData, len(src))
	for i, v := range src {
		rs[i] = NewDiskData(v)
	}
	return rs
}

func NewDiskData(src info.Disk) *DiskData {
	return &DiskData{
		Device:     src.Device(),
		MountPoint: src.MountPoint(),
		Type:       src.Type(),
		Total:      int(src.Total()),
		Used:       int(src.Used()),
		Available:  int(src.Available()),
	}
}

type DiskData struct {
	Device     string `json:"device"`
	MountPoint string `json:"mountpoint"`
	Type       string `json:"type"`
	Total      int    `json:"total"`
	Used       int    `json:"used"`
	Available  int    `json:"available"`
}

func NewLoadAverageData(src info.LoadAverage) *LoadAverageData {
	return &LoadAverageData{
		Load1:  src.Load1(),
		Load5:  src.Load5(),
		Load15: src.Load15(),
	}
}

type LoadAverageData struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func NewMemoryData() *MemoryData {
	return &MemoryData{}
}

type MemoryData struct {
	Swap    *SwapData    `json:"swap"`
	Virtual *VirtualData `json:"virtual"`
}

func NewSwapData(src info.Memory) *SwapData {
	return &SwapData{
		Total:     int(src.Total()),
		Used:      int(src.Used()),
		Available: int(src.Available()),
	}
}

type SwapData struct {
	Total     int `json:"total"`
	Used      int `json:"used"`
	Available int `json:"available"`
}

func NewVirtualData(src info.VirtualMemory) *VirtualData {
	return &VirtualData{
		Total:     int(src.Total()),
		Used:      int(src.Used()),
		Available: int(src.Available()),
		Cached:    int(src.Cached()),
	}
}

type VirtualData struct {
	Total     int `json:"total"`
	Used      int `json:"used"`
	Available int `json:"available"`
	Cached    int `json:"cached"`
}

func NewNetworkDataList(src []info.Network) []*NetworkData {
	rs := make([]*NetworkData, len(src))
	for i, v := range src {
		rs[i] = NewNetworkData(v)
	}
	return rs
}

func NewNetworkData(src info.Network) *NetworkData {
	return &NetworkData{
		Name:            src.Name(),
		HardwareAddress: src.HardwareAddress(),
		IPs:             src.IPs(),
	}
}

type NetworkData struct {
	Name            string   `json:"name"`
	HardwareAddress string   `json:"hardwareaddress"`
	IPs             []string `json:"ips"`
}

func NewTemperatureData(src info.Temperature) *TemperatureData {
	return &TemperatureData{
		Celsius:    src.Celsius(),
		Fahrenheit: src.Fahrenheit(),
	}
}

type TemperatureData struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

func NewTemperatureRoot() *TemperatureRoot {
	return &TemperatureRoot{}
}

type TemperatureRoot struct {
	CPU *TemperatureData `json:"cpu"`
}
