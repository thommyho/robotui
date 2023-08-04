package meter

import (
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/meter/tplink"
	"github.com/thommyho/robotui/util"
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
