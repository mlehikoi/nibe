package uplink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStatus(t *testing.T) {
	params := []Parameter{
		{Name: "40067", RawValue: 78},
		{Name: "40014", RawValue: 482},
		{Name: "40013", RawValue: 515},
		{Name: "40004", RawValue: 132},
		{Name: "40083", RawValue: 100},
		{Name: "40081", RawValue: 200},
		{Name: "40079", RawValue: 300},
	}
	status := newStatus(params)
	assert.InDelta(t, 7.8, status.AvgOutdoorTemp, .01)
	assert.InDelta(t, 48.2, status.HotWaterCharging, .01)
	assert.InDelta(t, 51.5, status.HotWaterTop, .01)
	assert.InDelta(t, 13.2, status.OutdoorTemp, .01)
	assert.Equal(t, 100, status.Current1)
	assert.Equal(t, 200, status.Current2)
	assert.Equal(t, 300, status.Current3)

}
