package core

import "github.com/thommyho/robotui/api"

// configProvider gives access to configuration repository
type configProvider interface {
	Meter(string) (api.Meter, error)
	Charger(string) (api.Charger, error)
	Vehicle(string) (api.Vehicle, error)
}
