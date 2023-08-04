package vehicle

import (
	"errors"
	"fmt"
	"time"

	"github.com/robotuimyhorobotuiotui/util"
	"github.com/robotuimyhorobotuiotui/vehicle/porsche"
	"github.com/samber/lo"
	"github.com/thommyho/robotui/api"
)

// Porsche is an api.Vehicle implementation for Porsche cars
type Porsche struct {
	*embed
	*porsche.Provider
}

func init() {
	registry.Add("porsche", NewPorscheFromConfig)
}

// NewPorscheFromConfig creates a new vehicle
func NewPorscheFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		embed               `mapstructure:",squash"`
		User, Password, VIN string
		Cache               time.Duration
	}{
		Cache: interval,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	log := util.NewLogger("porsche").Redact(cc.User, cc.Password, cc.VIN)
	identity := porsche.NewIdentity(log)

	ts, err := identity.Login(porsche.OAuth2Config, cc.User, cc.Password)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	api := porsche.NewAPI(log, ts)

	cc.VIN, err = ensureVehicle(cc.VIN, func() ([]string, error) {
		vehicles, err := api.Vehicles()
		return lo.Map(vehicles, func(v porsche.Vehicle, _ int) string {
			return v.VIN
		}), err
	})

	if err != nil {
		return nil, err
	}

	// check if vehicle is paired
	if res, err := api.PairingStatus(cc.VIN); err == nil && !porsche.IsPaired(res.Status) {
		return nil, errors.New("vehicle is not paired with the My Porsche account")
	}

	emobApi := porsche.NewEmobilityAPI(log, ts)
	capabilities, err := emobApi.Capabilities(cc.VIN)
	if err != nil {
		return nil, err
	}

	provider := porsche.NewProvider(log, api, emobApi, cc.VIN, capabilities.CarModel, cc.Cache)

	v := &Porsche{
		embed:    &cc.embed,
		Provider: provider,
	}

	return v, err
}
