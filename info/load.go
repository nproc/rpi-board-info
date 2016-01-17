package info

type LoadAverage interface {
	Load1() float64
	Load5() float64
	Load15() float64
}

type LoadAverageInformationProvider interface {
	LoadAverage() (LoadAverage, error)
}

func NewLoadAverage(load1, load5, load15 float64) LoadAverage {
	return &load{
		load1:  load1,
		load5:  load5,
		load15: load15,
	}
}

type load struct {
	load1  float64
	load5  float64
	load15 float64
}

func (l *load) Load1() float64 {
	return l.load1
}

func (l *load) Load5() float64 {
	return l.load5
}

func (l *load) Load15() float64 {
	return l.load15
}
