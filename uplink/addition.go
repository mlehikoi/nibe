package uplink

// Cat: ADDITION
// blocked:0
// imm. heater sensor:368
// fuse size:20
// time factor:7177
// electrical addition power:0
// set max electrical add.:560

type Addition struct {
	Blocked                 int     `name:"10033"`             // blocked        // blocked
	ImmHeaterSensor         float64 `name:"40024" scale:".1"`  // imm. heater sensor
	FuseSize                int     `name:"47214"`             // fuse size in amps
	TimeFactor              int     `name:"43081"`             // time factor
	ElectricalAdditionPower float64 `name:"43084" scale:".01"` // How much direct electrical heating is currently used in W?
	SetMaxElectricalAdd     float64 `name:"47212" scale:".01"` // User set maximum electrical addition power in kW
}

func newAddition(params []Parameter) *Addition {
	addition := &Addition{}
	parseParams(params, addition)
	return addition
}
