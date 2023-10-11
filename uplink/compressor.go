package uplink

// CPR_INFO_EP14
// blocked:0
// compressor starts:7406
// compressor sensor:330
// evaporator:59
// compressor operating time:10317
// compressor operating time hot water:2676

type Compressor struct {
	blocked                         int
	compressorStarts                int
	compressorSensor                int
	evaporator                      int
	compressorOperatingTime         int
	compressorOperatingTimeHotWater int
}

func newCompressor(params []Parameter) Compressor {
	m := newParamMap(params)
	return Compressor{
		blocked:                         m.int("blocked"),
		compressorStarts:                m.int("compressor starts"),
		compressorSensor:                m.int("compressor sensor"),
		evaporator:                      m.int("evaporator"),
		compressorOperatingTime:         m.int("compressor operating time"),
		compressorOperatingTimeHotWater: m.int("compressor operating time hot water"),
	}
}
