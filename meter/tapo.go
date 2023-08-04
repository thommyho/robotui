package meter

import (
	"github.com/robotuimyhorobotuiotui/meter/tapo"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/thommyho/robotui/api"
)

// TP-Link Tapo meter implementation
func init() {
	registry.Add("tapo", NewTapoFromConfig)
}

// NewTapoFromConfig creates a tapo meter from generic config
func NewTapoFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		URI      string
		User     string
		Password string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	return tapo.NewConnection(cc.URI, cc.User, cc.Password)
}
