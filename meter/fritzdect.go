package meter

import (
	"github.com/robotuimyhorobotuiotui/meter/fritzdect"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/thommyho/robotui/api"
)

// AVM FritzBox AHA interface specifications:
// https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AHA-HTTP-Interface.pdf

func init() {
	registry.Add("fritzdect", NewFritzDECTFromConfig)
}

// NewFritzDECTFromConfig creates a fritzdect meter from generic config
func NewFritzDECTFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc fritzdect.Settings
	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	return fritzdect.NewConnection(cc.URI, cc.AIN, cc.User, cc.Password)
}
