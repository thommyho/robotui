package vehicle

import (
	"time"

	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/request"
	"github.com/thommyho/robotui/vehicle/seat"
	"github.com/thommyho/robotui/vehicle/vag/service"
	"github.com/thommyho/robotui/vehicle/vag/tokenrefreshservice"
	"github.com/thommyho/robotui/vehicle/vw"
)

// https://github.com/trocotronic/weconnect
// https://github.com/TA2k/ioBroker.vw-connect

// Seat is an api.Vehicle implementation for Seat cars
type Seat struct {
	*embed
	*vw.Provider // provides the api implementations
}

func init() {
	registry.Add("seat", NewSeatFromConfig)
}

// NewSeatFromConfig creates a new vehicle
func NewSeatFromConfig(other map[string]interface{}) (api.Vehicle, error) {
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

	v := &Seat{
		embed: &cc.embed,
	}

	log := util.NewLogger("seat").Redact(cc.User, cc.Password, cc.VIN)

	trs := tokenrefreshservice.New(log, seat.TRSParams)
	ts, err := service.MbbTokenSource(log, trs, seat.AuthClientID, seat.AuthParams, cc.User, cc.Password)
	if err != nil {
		return nil, err
	}

	api := vw.NewAPI(log, ts, seat.Brand, seat.Country)
	api.Client.Timeout = cc.Timeout

	cc.VIN, err = ensureVehicle(cc.VIN, api.Vehicles)

	if err == nil {
		if err = api.HomeRegion(cc.VIN); err == nil {
			v.Provider = vw.NewProvider(api, cc.VIN, cc.Cache)
		}
	}

	return v, err
}
