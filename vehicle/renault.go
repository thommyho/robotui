package vehicle

import (
	"time"

	"github.com/robotuimyhorobotuiotui/util"
	"github.com/robotuimyhorobotuiotui/util/request"
	"github.com/robotuimyhorobotuiotui/vehicle/renault"
	"github.com/robotuimyhorobotuiotui/vehicle/renault/gigya"
	"github.com/robotuimyhorobotuiotui/vehicle/renault/kamereon"
	"github.com/robotuimyhorobotuiotui/vehicle/renault/keys"
	"github.com/thommyho/robotui/api"
)

// Credits to
//  https://github.com/hacf-fr/renault-api
//  https://github.com/edent/Renault-Zoe-API/issues/18
//  https://github.com/epenet/Renault-Zoe-API/blob/newapimockup/Test/MyRenault.py
//  https://github.com/jamesremuscat/pyze
//  https://muscatoxblog.blogspot.com/2019/07/delving-into-renaults-new-api.html

// Renault is an api.Vehicle implementation for Renault cars
type Renault struct {
	*embed
	*renault.Provider
}

func init() {
	registry.Add("dacia", func(other map[string]interface{}) (api.Vehicle, error) {
		return NewRenaultDaciaFromConfig("dacia", other)
	})
	registry.Add("renault", func(other map[string]interface{}) (api.Vehicle, error) {
		return NewRenaultDaciaFromConfig("renault", other)
	})
}

// NewRenaultDaciaFromConfig creates a new Renault/Dacia vehicle
func NewRenaultDaciaFromConfig(brand string, other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		embed                       `mapstructure:",squash"`
		User, Password, Region, VIN string
		Cache                       time.Duration
		Timeout                     time.Duration
	}{
		Region:  "de_DE",
		Cache:   interval,
		Timeout: request.Timeout,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	log := util.NewLogger(brand).Redact(cc.User, cc.Password, cc.VIN)

	v := &Renault{
		embed: &cc.embed,
	}

	keys := keys.New(log)
	keys.Load(cc.Region)

	identity := gigya.NewIdentity(log, keys.Gigya)
	if err := identity.Login(cc.User, cc.Password); err != nil {
		return nil, err
	}

	api := kamereon.New(log, keys.Kamereon, identity, func() error {
		return identity.Login(cc.User, cc.Password)
	})
	api.Client.Timeout = cc.Timeout

	accountID, err := api.Person(identity.PersonID, brand)

	if err != nil {
		return nil, err
	}

	vehicle, err := ensureVehicleEx(cc.VIN,
		func() ([]kamereon.Vehicle, error) {
			return api.Vehicles(accountID)
		},
		func(v kamereon.Vehicle) string {
			return v.VIN
		},
	)

	if err == nil {
		err = vehicle.Available()
	}

	v.Provider = renault.NewProvider(api, accountID, vehicle.VIN, cc.Cache)

	return v, err
}
