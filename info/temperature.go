package info

type Temperature interface {
	Celsius() float64
	Fahrenheit() float64
}

type TemperatureInformationProvider interface {
	CPUTemperature() (Temperature, error)
}

func NewTemperature(value int) Temperature {
	return &temperature{
		value: value,
	}
}

type temperature struct {
	value int
}

func (t *temperature) Celsius() float64 {
	return float64(t.value) / 1000
}

func (t *temperature) Fahrenheit() float64 {
	return (t.Celsius() * 1.8) + 32
}
