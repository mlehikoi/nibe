package uplink

// Cat: STATUS
// avg. outdoor temp:74
// hot water charging:482
// hot water top:515
// outdoor temp.:132
// current:0
// current:0
// current:0

type Status struct {
	OutdoorTemp      float64 `name:"40004" scale:".1"` // Current outdoor temperature in Celsius
	AvgOutdoorTemp   float64 `name:"40067" scale:".1"` // 24-hour average outdoor temperature in Celsius
	HotWaterCharging float64 `name:"40014" scale:".1"` // Hot water charging temperature in Celsius
	HotWaterTop      float64 `name:"40013" scale:".1"` // Hot water top temperature in Celsius
	Current1         int     `name:"40083"`            // Phase 1 current in Amps
	Current2         int     `name:"40081"`            // Phase 2 current in Amps
	Current3         int     `name:"40079"`            // Phase 3 current in Amps
}

func newStatus(params []Parameter) *Status {
	status := &Status{}
	parseParams(params, status)
	return status
}
