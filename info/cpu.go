package info

type CPU interface {
	Model() string
	Cores() int
}

type CPUInformationProvider interface {
	CPU() (CPU, error)
}

func NewCPU(model string, cores int) CPU {
	return &cpu{
		model: model,
		cores: cores,
	}
}

type cpu struct {
	model string
	cores int
}

func (c *cpu) Model() string {
	return c.model
}

func (c *cpu) Cores() int {
	return c.cores
}
