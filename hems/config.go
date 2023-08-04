package hems

import (
	"errors"
	"strings"

	"github.com/thommyho/robotui/core/site"
	"github.com/thommyho/robotui/hems/ocpp"
	"github.com/thommyho/robotui/hems/semp"
	"github.com/thommyho/robotui/server"
)

// HEMS describes the HEMS system interface
type HEMS interface {
	Run()
}

// NewFromConfig creates new HEMS from config
func NewFromConfig(typ string, other map[string]interface{}, site site.API, httpd *server.HTTPd) (HEMS, error) {
	switch strings.ToLower(typ) {
	case "sma", "shm", "semp":
		return semp.New(other, site, httpd)
	case "ocpp":
		return ocpp.New(other, site)
	default:
		return nil, errors.New("unknown hems: " + typ)
	}
}
