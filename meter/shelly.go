package meter

import (
	"github.com/robotuimyhorobotuiotui/meter/shelly"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/thommyho/robotui/api"
)

// Shelly meter implementation
func init() {
	registry.Add("shelly", NewShellyFromConfig)
}

// NewShellyFromConfig creates a Shelly charger from generic config
func NewShellyFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		URI      string
		User     string
		Password string
		Channel  int
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	conn, err := shelly.NewConnection(cc.URI, cc.User, cc.Password, cc.Channel)
	if err != nil {
		return nil, err
	}

	return shelly.NewSwitch(conn), nil
}
