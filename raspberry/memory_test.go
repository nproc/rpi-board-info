package raspberry_test

import (
	"testing"

	"github.com/nproc/rpi-board-info/raspberry"
	"github.com/stretchr/testify/assert"
)

var m raspberry.Memory

func init() {
	m = raspberry.Memory{
		Total:     8252903424,
		Used:      3724775424,
		Cached:    1384595456,
		Available: 6032896000,
	}
}

func TestMemoryToAvailablePercent(t *testing.T) {
	assert.Equal(t, m.ToAvailablePercent(), 73)
}

func TestMemoryToUsedPercent(t *testing.T) {
	assert.Equal(t, m.ToUsedPercent(), 27)
}
