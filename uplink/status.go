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
	OutdoorTemp      float32 // Current outdoor temperature in Celsius
	AvgOutdoorTemp   float32 // 24-hour average outdoor temperature in Celsius
	HotWaterCharging float32 // Hot water charging temperature in Celsius
	HotWaterTop      float32 // Hot water top temperature in Celsius
	Current1         int     // Phase 1 current in Amps
	Current2         int     // Phase 2 current in Amps
	Current3         int     // Phase 3 current in Amps
}

func newStatus(params []Parameter) Status {
	m := newParamMap(params)
	return Status{
		AvgOutdoorTemp:   m.float32("avg. outdoor temp") / 10.,
		HotWaterCharging: m.float32("hot water charging") / 10.,
		HotWaterTop:      m.float32("hot water top") / 10.,
		OutdoorTemp:      m.float32("outdoor temp.") / 10.,
		Current1:         m.int("current", 0),
		Current2:         m.int("current", 1),
		Current3:         m.int("current", 2),
	}
}
