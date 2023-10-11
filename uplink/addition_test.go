package uplink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAddition(t *testing.T) {
	params := []Parameter{
		{Name: "10033", RawValue: 0},
		{Name: "40024", RawValue: 396},
		{Name: "47214", RawValue: 20},
		{Name: "43081", RawValue: 7185},
		{Name: "43084", RawValue: 100},
		{Name: "47212", RawValue: 560},
	}
	addition := newAddition(params)
	assert.Equal(t, 0, addition.Blocked)
	assert.InDelta(t, 39.6, addition.ImmHeaterSensor, .01)
	assert.Equal(t, 20, addition.FuseSize)
	assert.Equal(t, 7185, addition.TimeFactor)
	assert.InDelta(t, 1.0, addition.ElectricalAdditionPower, .01)
	assert.InDelta(t, 5.6, addition.SetMaxElectricalAdd, .01)

}
