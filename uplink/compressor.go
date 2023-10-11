package uplink

// CPR_INFO_EP14
// blocked:0
// compressor starts:7406
// compressor sensor:330
// evaporator:59
// compressor operating time:10317
// compressor operating time hot water:2676

type Compressor struct {
	Blocked                         int `name:"10012" scale:"1.0"`
	CompressorStarts                int `name:"43416" scale:"1.0"`
	CompressorSensor                int `name:"40023" scale:"1.0"`
	Evaporator                      int `name:"40020" scale:"1.0"`
	CompressorOperatingTime         int `name:"43420" scale:"1.0"`
	CompressorOperatingTimeHotWater int `name:"43424" scale:"1.0"`
}

func newCompressor(params []Parameter) *Compressor {
	compressor := &Compressor{}
	parseParams(params, compressor)
	return compressor
}
