package uplink

import "time"

type System struct {
	systemId         uint
	Name             string
	ProductName      string
	SecurityLevel    string
	SerialNumber     string
	LastActivityDate time.Time
	ConnectionStatus string
	HasAlarmed       bool

	Status      *Status      // Basic status information
	Compressor  *Compressor  // Compressor information
	Ventilation *Ventilation // Ventilation information
	Climate     *Climate     // Climate system information
	Addition    *Addition    // Additional electric heating
}

func newSystem(obj Object) System {
	return System{
		systemId:         obj.SystemId,
		Name:             obj.Name,
		ProductName:      obj.ProductName,
		SecurityLevel:    obj.SecurityLevel,
		SerialNumber:     obj.SerialNumber,
		LastActivityDate: obj.LastActivityDate,
		ConnectionStatus: obj.ConnectionStatus,
		HasAlarmed:       obj.HasAlarmed,
	}
}
