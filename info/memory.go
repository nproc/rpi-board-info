package info

type Memory interface {
	Total() uint64
	Used() uint64
	Available() uint64
}

type VirtualMemory interface {
	Memory
	Cached() uint64
}

type MemoryInformationProvider interface {
	VirtualMemory() (VirtualMemory, error)
	SwapMemory() (Memory, error)
}

func NewMemory(total, used, available uint64) Memory {
	return &memory{
		total:     total,
		used:      used,
		available: available,
	}
}

func NewVirtualMemory(total, used, available, cached uint64) VirtualMemory {
	return &virtualMemory{
		memory: memory{
			total:     total,
			used:      used,
			available: available,
		},
		cached: cached,
	}
}

type memory struct {
	total     uint64
	used      uint64
	available uint64
}

func (m *memory) Total() uint64 {
	return m.total
}

func (m *memory) Used() uint64 {
	return m.used
}

func (m *memory) Available() uint64 {
	return m.available
}

type virtualMemory struct {
	memory
	cached uint64
}

func (m *virtualMemory) Cached() uint64 {
	return m.cached
}
