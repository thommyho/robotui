package ocpp

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/site"
	"github.com/evcc-io/evcc/hems/ocpp/profile"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/machine"

	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	ocppcore "github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ws"
)

// OCPP is an OCPP client
type OCPP struct {
	log  *util.Logger
	site site.API
	cp   ocpp16.ChargePoint
}

const retryTimeout = 5 * time.Second

// New generates OCPP chargepoint client
func New(conf map[string]interface{}, site site.API) (*OCPP, error) {
	var cc struct {
		URI       string
		StationID string
	}

	if err := util.DecodeOther(conf, &cc); err != nil {
		return nil, err
	}

	log := util.NewLogger("ocpp")

	if cc.StationID == "" {
		id, err := machine.ProtectedID("evcc-ocpp")
		if err == nil {
			cc.StationID = fmt.Sprintf("evcc-%s", strings.ToLower(id))
		} else {
			cc.StationID = fmt.Sprintf("evcc-%d", rand.Int31())
		}
		log.DEBUG.Println("station id:", cc.StationID)
	}

	ws := ws.NewClient()
	cp := ocpp16.NewChargePoint(cc.StationID, nil, ws)

	s := &OCPP{
		log:  log,
		site: site,
		cp:   cp,
	}

	err := cp.Start(cc.URI)
	if err == nil {
		cp.SetCoreHandler(profile.NewCore(log, profile.GetDefaultConfig()))
		cp.SetSmartChargingHandler(profile.NewSmartCharging(log))

		go s.errorHandler(ws.Errors())
		go s.errorHandler(cp.Errors())
	}

	return s, err
}

// errorHandler logs error channel
func (s *OCPP) errorHandler(errC <-chan error) {
	for err := range errC {
		s.log.ERROR.Println(err)
	}
}

// Run executes the OCPP chargepoint client
func (s *OCPP) Run() {
	for {
		for id, lp := range s.site.Loadpoints() {
			connector := id + 1

			status := ocppcore.ChargePointStatusAvailable
			if lp.GetStatus() == api.StatusC {
				status = ocppcore.ChargePointStatusCharging
			}

			s.log.DEBUG.Printf("send: lp-%d status: %+v", connector, status)
			if _, err := s.cp.StatusNotification(connector, ocppcore.NoError, status); err != nil {
				s.log.ERROR.Printf("lp-%d: %v", connector, err)
			}
		}

		time.Sleep(retryTimeout)
	}
}
