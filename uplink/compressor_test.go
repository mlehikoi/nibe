package uplink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCompressor(t *testing.T) {
	params := []Parameter{
		{Name: "40079", RawValue: 0},
		{Name: "10012", RawValue: 0},
		{Name: "43416", RawValue: 7406},
		{Name: "40023", RawValue: 320},
		{Name: "40020", RawValue: 189},
		{Name: "43420", RawValue: 10318},
		{Name: "43424", RawValue: 2676},
	}
	compressor := newCompressor(params)
	assert.Equal(t, 0, compressor.Blocked)
	assert.Equal(t, 7406, compressor.CompressorStarts)
	assert.Equal(t, 320, compressor.CompressorSensor)
	assert.Equal(t, 189, compressor.Evaporator)
	assert.Equal(t, 10318, compressor.CompressorOperatingTime)
	assert.Equal(t, 2676, compressor.CompressorOperatingTimeHotWater)
}
