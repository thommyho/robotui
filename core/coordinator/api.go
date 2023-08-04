package coordinator

import "github.com/thommyho/robotui/api"

// API is the coordinator API
type API interface {
	GetVehicles() []api.Vehicle
	Acquire(api.Vehicle)
	Release(api.Vehicle)
	IdentifyVehicleByStatus() api.Vehicle
	GetVehicleIndex(api.Vehicle) int
}
