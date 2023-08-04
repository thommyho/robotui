package vehicle

import (
	"context"
	"time"

	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/request"
	"github.com/thommyho/robotui/vehicle/audi/etron"
	"github.com/thommyho/robotui/vehicle/vag/idkproxy"
	"github.com/thommyho/robotui/vehicle/vag/service"
	"github.com/thommyho/robotui/vehicle/vag/vwidentity"
	"github.com/thommyho/robotui/vehicle/vw/id"
)

// https://github.com/TA2k/ioBroker.vw-connect
// https://github.com/arjenvrh/audi_connect_ha/blob/master/custom_components/audiconnect/audi_services.py

// Audi is an api.Vehicle implementation for Audi cars
type Audi struct {
	*embed
	*id.Provider // provides the api implementations
}

func init() {
	registry.Add("audi", NewAudiFromConfig)
	registry.Add("etron", NewAudiFromConfig)
}

// NewAudiFromConfig creates a new vehicle
func NewAudiFromConfig(other map[string]interface{}) (api.Vehicle, error) {
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

	v := &Audi{
		embed: &cc.embed,
	}

	log := util.NewLogger("audi").Redact(cc.User, cc.Password, cc.VIN)

	// get initial VW identity id_token
	q, err := vwidentity.Login(log, etron.AuthParams, cc.User, cc.Password)
	if err != nil {
		return nil, err
	}

	// exchange initial VW identity id_token for Audi AAZS token
	idk := idkproxy.New(log, etron.IDKParams)
	ats, its, err := service.AAZSTokenSource(log, idk, etron.AZSConfig, q)
	if err != nil {
		return nil, err
	}

	// use the etron API for list of vehicles
	api := etron.NewAPI(log, ats)

	vehicle, err := ensureVehicleEx(
		cc.VIN, func() ([]etron.Vehicle, error) {
			ctx, cancel := context.WithTimeout(context.Background(), cc.Timeout)
			defer cancel()
			return api.Vehicles(ctx)
		},
		func(v etron.Vehicle) string {
			return v.VIN
		},
	)

	if err == nil {
		if v.Title_ == "" {
			v.Title_ = vehicle.Nickname
		}

		api := id.NewAPI(log, its)
		api.Client.Timeout = cc.Timeout

		v.Provider = id.NewProvider(api, vehicle.VIN, cc.Cache)
	}

	return v, err
}
