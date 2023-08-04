package meter

import (
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/meter/tasmota"
	"github.com/thommyho/robotui/util"
)

// Tasmota meter implementation
type Tasmota struct {
	conn  *tasmota.Connection
	usage string
}

// Tasmota meter implementation
func init() {
	registry.Add("tasmota", NewTasmotaFromConfig)
}

// NewTasmotaFromConfig creates a Tasmota meter from generic config
func NewTasmotaFromConfig(other map[string]interface{}) (api.Meter, error) {
	cc := struct {
		URI      string
		User     string
		Password string
		Channel  int
		Usage    string
	}{
		Channel: 1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewTasmota(cc.URI, cc.User, cc.Password, cc.Usage, cc.Channel)
}

// NewTasmota creates Tasmota meter
func NewTasmota(uri, user, password, usage string, channel int) (*Tasmota, error) {
	conn, err := tasmota.NewConnection(uri, user, password, channel)
	if err != nil {
		return nil, err
	}

	c := &Tasmota{
		conn:  conn,
		usage: usage,
	}

	return c, err
}

var _ api.Meter = (*Tasmota)(nil)

// CurrentPower implements the api.Meter interface
func (c *Tasmota) CurrentPower() (float64, error) {
	if c.usage == "grid" {
		return c.conn.SmlPower()
	}
	return c.conn.CurrentPower()
}

var _ api.MeterEnergy = (*Tasmota)(nil)

// TotalEnergy implements the api.MeterEnergy interface
func (c *Tasmota) TotalEnergy() (float64, error) {
	if c.usage == "grid" {
		return c.conn.SmlTotalEnergy()
	}
	return c.conn.TotalEnergy()
}
