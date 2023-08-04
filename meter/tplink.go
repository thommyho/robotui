package meter

import (
	"github.com/robotuimyhorobotuiotui/meter/tplink"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/thommyho/robotui/api"
)

func init() {
	registry.Add("tplink", NewTPLinkFromConfig)
}

// NewTPLinkFromConfig creates a tapo meter from generic config
func NewTPLinkFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		URI string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return tplink.NewConnection(cc.URI)
}
