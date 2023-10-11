package uplink

// Cat: VENTILATION
// fan speed:62
// exhaust air:205
// extract air:25
// supply air:137

type Ventilation struct {
	FanSpeed    int     `name:"10001"`            // Fan speed in percentage
	ExhaustTemp float64 `name:"40025" scale:".1"` // exhaust air temperature, before compressor
	ExtractTemp float64 `name:"40026" scale:".1"` // extract air, after compressor
	SupplyTemp  float64 `name:"40075" scale:".1"` // intake air temperature

}

func newVentilation(params []Parameter) *Ventilation {
	vent := &Ventilation{}
	parseParams(params, vent)
	return vent
}
