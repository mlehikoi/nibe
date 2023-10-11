package uplink

// Cat: ADDITION
// blocked:0
// imm. heater sensor:368
// fuse size:20
// time factor:7177
// electrical addition power:0
// set max electrical add.:560

type Addition struct {
	Blocked                  int     // blocked
	ImmersedHeaterSensor     float32 // imm. heater sensor
	FuseSize                 int     // fuse size in amps
	TimeFactor               int     // time factor
	ElectricalAdditionPower  float32 // How much direct electrical heating is currently used in W?
	SetMaxElectricalAddition float32 // User set maximum electrical addition power in kW
}

func newAddition(params []Parameter) Addition {
	m := newParamMap(params)
	return Addition{
		Blocked:                  m.int("blocked"),
		ImmersedHeaterSensor:     m.float32("imm. heater sensor") / 10.,
		FuseSize:                 m.int("fuse size"),
		TimeFactor:               m.int("time factor"),
		ElectricalAdditionPower:  m.float32("electrical addition power") / 100.,
		SetMaxElectricalAddition: m.float32("set max electrical add.") / 100.,
	}
}
