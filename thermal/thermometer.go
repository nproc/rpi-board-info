package thermal

import (
	"os/exec"
	"strconv"
	"strings"
)

// Thermometer provides a simple API to measure the current
// system temperature
type Thermometer struct {
}

// Measure measures the current system temperature
// https://www.kernel.org/doc/Documentation/thermal/sysfs-api.txt
func (t Thermometer) Measure() Temperature {
	output, err := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()

	if err != nil {
		return Temperature(0)
	}

	celsius, err := strconv.Atoi(strings.TrimSpace(string(output)))

	if err != nil {
		return Temperature(0)
	}

	return TemperatureFromCelsiusMillidegree(celsius)
}
