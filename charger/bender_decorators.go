package charger

// Code generated by github.com/thommyho/robotui/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/thommyho/robotui/api"
)

func decorateBenderCC(base *BenderCC, meter func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error), chargeRater func() (float64, error), meterEnergy func() (float64, error), identifier func() (string, error)) api.Charger {
	switch {
	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return base

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Meter
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.PhaseCurrents
		}{
			BenderCC: base,
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Meter
			api.PhaseCurrents
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.PhaseVoltages
		}{
			BenderCC: base,
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Meter
			api.PhaseVoltages
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.MeterEnergy
		}{
			BenderCC: base,
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Meter
			api.MeterEnergy
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.MeterEnergy
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.MeterEnergy
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.PhaseCurrents
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.MeterEnergy
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.MeterEnergy
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargeRater != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*BenderCC
			api.ChargeRater
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			BenderCC: base,
			ChargeRater: &decorateBenderCCChargeRaterImpl{
				chargeRater: chargeRater,
			},
			Identifier: &decorateBenderCCIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateBenderCCMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateBenderCCMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateBenderCCPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateBenderCCPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decorateBenderCCChargeRaterImpl struct {
	chargeRater func() (float64, error)
}

func (impl *decorateBenderCCChargeRaterImpl) ChargedEnergy() (float64, error) {
	return impl.chargeRater()
}

type decorateBenderCCIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateBenderCCIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateBenderCCMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateBenderCCMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateBenderCCMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateBenderCCMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateBenderCCPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateBenderCCPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateBenderCCPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateBenderCCPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
