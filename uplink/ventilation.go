package uplink

// Cat: VENTILATION
// fan speed:62
// exhaust air:205
// extract air:25
// supply air:137

type Ventilation struct {
	FanSpeed    int     // Fan speed in percentage
	exhaustTemp float32 // exhaust air temperature, before compressor
	extractTemp float32 // extract air, after compressor
	supplyTemp  float32 // intake air temperature
}

func newVentilation(params []Parameter) Ventilation {
	m := newParamMap(params)
	return Ventilation{
		FanSpeed:    m.int("fan speed"),
		exhaustTemp: m.float32("exhaust air") / 10.,
		extractTemp: m.float32("extract air") / 10.,
		supplyTemp:  m.float32("supply air") / 10.,
	}
}
