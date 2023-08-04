package charger

// Code generated by github.com/thommyho/robotui/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/robotuimyhorobotuiotui/api"
)

func decoratePhoenixEMEth(base *PhoenixEMEth, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterEnergy == nil && phaseCurrents == nil:
		return base

	case meter != nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*PhoenixEMEth
			api.MeterEnergy
		}{
			PhoenixEMEth: base,
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*PhoenixEMEth
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter == nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*PhoenixEMEth
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEMEthMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decoratePhoenixEMEthMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEMEthMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEMEthMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixEMEthMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decoratePhoenixEMEthPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEMEthPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}
