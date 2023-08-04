package vehicle

import (
	"time"

	"github.com/robotuimyhorobotuiotui/util"
	"github.com/robotuimyhorobotuiotui/util/request"
	"github.com/robotuimyhorobotuiotui/vehicle/vag/loginapps"
	"github.com/robotuimyhorobotuiotui/vehicle/vag/vwidentity"
	"github.com/robotuimyhorobotuiotui/vehicle/vw/id"
	"github.com/thommyho/robotui/api"
)

// https://github.com/TA2k/ioBroker.vw-connect

// ID is an api.Vehicle implementation for ID cars
type ID struct {
	*embed
	*id.Provider // provides the api implementations
}

func init() {
	registry.Add("id", NewIDFromConfig)
}

// NewIDFromConfig creates a new vehicle
func NewIDFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		embed               `mapstructure:",squash"`
		User, Password, VIN string
		Cache               time.Duration
		Timeout             time.Duration
	}{
		Cache:   interval,
		Timeout: request.Timeout,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	v := &ID{
		embed: &cc.embed,
	}

	log := util.NewLogger("id").Redact(cc.User, cc.Password, cc.VIN)

	q, err := vwidentity.LoginWithAuthURL(log, id.LoginURL, id.AuthParams, cc.User, cc.Password)
	if err != nil {
		return nil, err
	}

	apps := loginapps.New(log)
	token, err := apps.Exchange(q)
	if err != nil {
		return nil, err
	}

	api := id.NewAPI(log, apps.TokenSource(token))
	api.Client.Timeout = cc.Timeout

	vehicle, err := ensureVehicleEx(
		cc.VIN, api.Vehicles,
		func(v id.Vehicle) string {
			return v.VIN
		},
	)

	if v.Title_ == "" {
		v.Title_ = vehicle.Nickname
	}

	if err == nil {
		v.Provider = id.NewProvider(api, vehicle.VIN, cc.Cache)
	}

	return v, err
}
