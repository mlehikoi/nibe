package uplink

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseClimate(t *testing.T) {
	params := []Parameter{
		{Name: "43161", RawValue: 0},
		{Name: "43009", RawValue: 300},
		{Name: "40071", RawValue: -32768},
		{Name: "40008", RawValue: 304},
		{Name: "40012", RawValue: 283},
		{Name: "40033", RawValue: -32768},
	}
	climate := newClimate(params)
	assert.Equal(t, 0, climate.ExternalAdjustment)
	assert.InDelta(t, 30, climate.CalculatedFlowTemp, .01)
	assert.True(t, math.IsNaN(climate.ExternalFlowTemp))
	assert.InDelta(t, 30.4, climate.HeatMediumFlow, .01)
	assert.InDelta(t, 28.3, climate.ReturnTemp, .01)
	assert.True(t, math.IsNaN(climate.RoomTemp), .01)

}
