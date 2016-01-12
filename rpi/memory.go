package rpi

import "math"

// Memory represents the system memory information
type Memory struct {
	Total     uint64
	Used      uint64
	Cached    uint64
	Available uint64
}

// ToUsedPercent returns how many percent of the total system
// memory is in use
func (m Memory) ToUsedPercent() int {
	percent := float64(m.Total-m.Available) / float64(m.Total) * 100.0

	return int(math.Floor(percent + .5))
}

// ToAvailablePercent returns how many percent of the total
// system memory is still available
func (m Memory) ToAvailablePercent() int {
	percent := float64(m.Available) / float64(m.Total) * 100.0

	return int(math.Floor(percent + .5))
}
