package uplink

// Cat: SYSTEM_1
// external adjustment:0
// calculated flow temp.:270
// external flow temp.:-32768
// heat medium flow:274
// return temp.:252
// room temperature:-32768

type Climate struct {
	ExternalAdjustment int     // External adjustment
	CalculatedFlowTemp float32 // Flow target temp
	ExternalFlowTemp   float32 // External flow temp, -3276.8 if not available
	HeatMediumFlow     float32 // Flow temp at medium point
	ReturnFlowTemp     float32 // Flow return temp
	RoomTemp           float32 // Room temp:, -3276,8 if not available
}

func newClimate(params []Parameter) Climate {
	m := newParamMap(params)
	return Climate{
		ExternalAdjustment: m.int("external adjustment"),
		CalculatedFlowTemp: m.float32("calculated flow temp.") / 10.,
		ExternalFlowTemp:   m.float32("external flow temp.") / 10.,
		HeatMediumFlow:     m.float32("heat medium flow") / 10.,
		ReturnFlowTemp:     m.float32("return temp.") / 10.,
		RoomTemp:           m.float32("room temperature") / 10.,
	}
}
