package modbus

import (
	"fmt"
	"net"

	"github.com/andig/mbserver"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/robotuimyhorobotuiotui/util/modbus"
	"github.com/robotuimyhorobotuiotui/util/sponsor"
	"github.com/thommyho/robotui/api"
)

func StartProxy(port int, config modbus.Settings, readOnly bool) error {
	conn, err := modbus.NewConnection(config.URI, config.Device, config.Comset, config.Baudrate, modbus.ProtocolFromRTU(config.RTU), config.ID)
	if err != nil {
		return err
	}

	if !sponsor.IsAuthorized() {
		return api.ErrSponsorRequired
	}

	h := &handler{
		log:            util.NewLogger(fmt.Sprintf("proxy-%d", port)),
		readOnly:       readOnly,
		RequestHandler: new(mbserver.DummyHandler), // supplies HandleDiscreteInputs
		conn:           conn,
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	h.log.DEBUG.Printf("modbus proxy for %s listening at :%d", config.String(), port)

	srv, err := mbserver.New(h, mbserver.Logger(&logger{log: h.log}))

	if err == nil {
		err = srv.Start(l)
	}

	return err
}
