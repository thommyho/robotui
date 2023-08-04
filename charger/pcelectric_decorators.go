package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePCE(base *PCElectric, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterEnergy == nil && phaseCurrents == nil:
		return base

	case meter != nil && meterEnergy == nil && phaseCurrents == nil:
		return &struct {
			*PCElectric
			api.Meter
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*PCElectric
			api.MeterEnergy
		}{
			PCElectric: base,
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil:
		return &struct {
			*PCElectric
			api.Meter
			api.MeterEnergy
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*PCElectric
			api.PhaseCurrents
		}{
			PCElectric: base,
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil:
		return &struct {
			*PCElectric
			api.Meter
			api.PhaseCurrents
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter == nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*PCElectric
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PCElectric: base,
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil:
		return &struct {
			*PCElectric
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decoratePCEMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePCEMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePCEMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePCEMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decoratePCEPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePCEPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}