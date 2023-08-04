package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateEEBus(base *EEBus, meter func() (float64, error), phaseCurrents func() (float64, float64, float64, error), chargeRater func() (float64, error)) api.Charger {
	switch {
	case chargeRater == nil && meter == nil && phaseCurrents == nil:
		return base

	case chargeRater == nil && meter != nil && phaseCurrents == nil:
		return &struct {
			*EEBus
			api.Meter
		}{
			EEBus: base,
			Meter: &decorateEEBusMeterImpl{
				meter: meter,
			},
		}

	case chargeRater == nil && meter == nil && phaseCurrents != nil:
		return &struct {
			*EEBus
			api.PhaseCurrents
		}{
			EEBus: base,
			PhaseCurrents: &decorateEEBusPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && meter != nil && phaseCurrents != nil:
		return &struct {
			*EEBus
			api.Meter
			api.PhaseCurrents
		}{
			EEBus: base,
			Meter: &decorateEEBusMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEEBusPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && meter == nil && phaseCurrents == nil:
		return &struct {
			*EEBus
			api.ChargeRater
		}{
			EEBus: base,
			ChargeRater: &decorateEEBusChargeRaterImpl{
				chargeRater: chargeRater,
			},
		}

	case chargeRater != nil && meter != nil && phaseCurrents == nil:
		return &struct {
			*EEBus
			api.ChargeRater
			api.Meter
		}{
			EEBus: base,
			ChargeRater: &decorateEEBusChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateEEBusMeterImpl{
				meter: meter,
			},
		}

	case chargeRater != nil && meter == nil && phaseCurrents != nil:
		return &struct {
			*EEBus
			api.ChargeRater
			api.PhaseCurrents
		}{
			EEBus: base,
			ChargeRater: &decorateEEBusChargeRaterImpl{
				chargeRater: chargeRater,
			},
			PhaseCurrents: &decorateEEBusPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && meter != nil && phaseCurrents != nil:
		return &struct {
			*EEBus
			api.ChargeRater
			api.Meter
			api.PhaseCurrents
		}{
			EEBus: base,
			ChargeRater: &decorateEEBusChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateEEBusMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEEBusPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decorateEEBusChargeRaterImpl struct {
	chargeRater func() (float64, error)
}

func (impl *decorateEEBusChargeRaterImpl) ChargedEnergy() (float64, error) {
	return impl.chargeRater()
}

type decorateEEBusMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateEEBusMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateEEBusPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateEEBusPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}