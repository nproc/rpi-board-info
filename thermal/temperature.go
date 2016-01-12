package thermal

// TemperatureFromCelsiusMillidegree creates a temperature representation
// from celsius millidegrees
func TemperatureFromCelsiusMillidegree(celsius int) Temperature {
	return Temperature(celsius)
}

// Temperature defines a generic type for representing a temperature
type Temperature int

// ToCelsius returns the current temperature in celsius
func (t Temperature) ToCelsius() float64 {
	return float64(t) / 1000
}

// ToFahrenheit returns the current temperature in fahrenheit
func (t Temperature) ToFahrenheit() float64 {
	return (float64(t) / 1000 * 1.8) + 32
}
