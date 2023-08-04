package site

import (
	"github.com/robotuimyhorobotuiotui/core/loadpoint"
	"github.com/thommyho/robotui/api"
)

// API is the external site API
type API interface {
	Healthy() bool
	Loadpoints() []loadpoint.API

	//
	// battery
	//

	GetBufferSoc() float64
	SetBufferSoc(float64) error
	GetBufferStartSoc() float64
	SetBufferStartSoc(float64) error
	GetPrioritySoc() float64
	SetPrioritySoc(float64) error

	//
	// power and energy
	//

	GetResidualPower() float64
	SetResidualPower(float64) error

	//
	// vehicles
	//

	// GetVehicles is the list of vehicles
	GetVehicles() []api.Vehicle

	//
	// tariffs and costs
	//

	// GetTariff returns the respective tariff
	GetTariff(string) api.Tariff
	GetSmartCostLimit() float64
	SetSmartCostLimit(float64) error
}
