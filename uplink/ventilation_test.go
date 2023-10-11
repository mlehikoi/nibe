package uplink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVentilation(t *testing.T) {
	params := []Parameter{
		{Name: "10001", RawValue: 62},
		{Name: "40025", RawValue: 228},
		{Name: "40026", RawValue: 175},
		{Name: "40075", RawValue: 222},
	}
	ventilation := newVentilation(params)
	assert.Equal(t, 62, ventilation.FanSpeed)
	assert.InDelta(t, 22.8, ventilation.ExhaustTemp, .01)
	assert.InDelta(t, 17.5, ventilation.ExtractTemp, .01)
	assert.InDelta(t, 22.2, ventilation.SupplyTemp, .01)

}
