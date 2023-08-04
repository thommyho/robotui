package charger

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/remotetrigger"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
	"github.com/stretchr/testify/suite"
	"github.com/thommyho/robotui/charger/ocpp"
)

const (
	ocppTestUrl            = "ws://localhost:8887"
	ocppTestConnectTimeout = 10 * time.Second
	ocppTestTimeout        = 3 * time.Second
	ocppTestConnector      = 1
)

func TestOcpp(t *testing.T) {
	suite.Run(t, new(ocppTestSuite))
}

type ocppTestSuite struct {
	suite.Suite
	clock *clock.Mock
	cp    ocpp16.ChargePoint
}

func (suite *ocppTestSuite) SetupSuite() {
	// setup cs
	suite.NotNil(ocpp.Instance())

	// setup cp
	suite.clock = clock.NewMock()
	cp := ocpp16.NewChargePoint("test", nil, nil)

	// set a handler for all callback functions
	triggerC := make(chan remotetrigger.MessageTrigger, 1)
	handler := &ChargePointHandler{triggerC: triggerC}
	cp.SetCoreHandler(handler)
	cp.SetRemoteTriggerHandler(handler)

	go func() {
		for msg := range triggerC {
			suite.handleTrigger(msg)
		}
	}()

	suite.cp = cp
}

func (suite *ocppTestSuite) handleTrigger(msg remotetrigger.MessageTrigger) {
	switch msg {
	case core.BootNotificationFeatureName:
		if res, err := suite.cp.BootNotification("demo", "robotui"); err != nil {
			suite.T().Log("BootNotification:", err)
		} else {
			suite.T().Log("BootNotification:", res)
		}

	case core.StatusNotificationFeatureName:
		if res, err := suite.cp.StatusNotification(ocppTestConnector, core.NoError, core.ChargePointStatusAvailable); err != nil {
			suite.T().Log("StatusNotification:", err)
		} else {
			suite.T().Log("StatusNotification:", res)
		}

	case core.MeterValuesFeatureName:
		if res, err := suite.cp.MeterValues(1, []types.MeterValue{
			{
				Timestamp: types.NewDateTime(suite.clock.Now()),
				SampledValue: []types.SampledValue{
					{Measurand: types.MeasurandPowerActiveImport, Value: "1000"},
					{Measurand: types.MeasurandEnergyActiveImportRegister, Value: "1.2", Unit: "kWh"},
				},
			},
		}); err != nil {
			suite.T().Log("MeterValues:", err)
		} else {
			suite.T().Log("MeterValues:", res)
		}

	default:
		suite.T().Log(msg)
	}
}

func (suite *ocppTestSuite) TestConnect() {
	// start cp client
	suite.NoError(suite.cp.Start(ocppTestUrl))
	suite.True(suite.cp.IsConnected())

	// start cp server
	c, err := NewOCPP("test", ocppTestConnector, "", "", 0, false, false, ocppTestConnectTimeout, ocppTestTimeout)
	suite.NoError(err)

	if err != nil {
		return
	}

	suite.clock.Add(ocppTestTimeout)
	c.cp.TestClock(suite.clock)

	_, err = c.Status()
	suite.NoError(err)

	// power
	f, err := c.currentPower()
	suite.NoError(err)
	suite.Equal(1e3, f)

	// energy
	f, err = c.totalEnergy()
	suite.NoError(err)
	suite.Equal(1.2, f)
}
