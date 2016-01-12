package thermal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperatureFromCelsiusMillidegree(t *testing.T) {
	assert.NotNil(t, TemperatureFromCelsiusMillidegree(44000))
}

func TestToCelsius(t *testing.T) {
	assert.Equal(
		t,
		TemperatureFromCelsiusMillidegree(44000).ToCelsius(),
		44.0,
	)
}

func TestToFahrenheit(t *testing.T) {
	assert.Equal(
		t,
		TemperatureFromCelsiusMillidegree(44000).ToFahrenheit(),
		111.2,
	)
}
