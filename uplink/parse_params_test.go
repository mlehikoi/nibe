package uplink

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotANumber(t *testing.T) {
	params := []Parameter{
		{Name: "40067", RawValue: -32767},
		{Name: "40014", RawValue: -32768},
	}
	status := newStatus(params)
	assert.False(t, math.IsNaN(status.AvgOutdoorTemp))
	assert.True(t, math.IsNaN(status.HotWaterCharging))
}
