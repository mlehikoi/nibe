package uplink

// Cat: SYSTEM_1
// external adjustment:0
// calculated flow temp.:270
// external flow temp.:-32768
// heat medium flow:274
// return temp.:252
// room temperature:-32768

type Climate struct {
	ExternalAdjustment int     `name:"43161"`            // External adjustment
	CalculatedFlowTemp float64 `name:"43009" scale:".1"` // Flow target temp
	ExternalFlowTemp   float64 `name:"40071" scale:".1"` // External flow temp, -3276.8 if not available
	HeatMediumFlow     float64 `name:"40008" scale:".1"` // Flow temp at medium point
	ReturnTemp         float64 `name:"40012" scale:".1"` // Flow return temp
	RoomTemp           float64 `name:"40033" scale:".1"` // Room temp:, -3276,8 if not available
}

func newClimate(params []Parameter) *Climate {
	climate := &Climate{}
	parseParams(params, climate)
	return climate
}
