package vehicle

// Code generated by github.com/thommyho/robotui/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/robotuimyhorobotuiotui/api"
)

func decorateVehicle(base api.Vehicle, chargeState func() (api.ChargeStatus, error), vehicleRange func() (int64, error), vehicleOdometer func() (float64, error), vehicleClimater func() (bool, error), resurrector func() error) api.Vehicle {
	switch {
	case chargeState == nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange == nil:
		return base

	case chargeState != nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.VehicleRange
		}{
			Vehicle: base,
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.VehicleOdometer
		}{
			Vehicle: base,
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleOdometer
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.VehicleClimater
		}{
			Vehicle: base,
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleClimater
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.VehicleClimater
			api.VehicleRange
		}{
			Vehicle: base,
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleClimater
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.VehicleClimater
			api.VehicleOdometer
		}{
			Vehicle: base,
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleClimater
			api.VehicleOdometer
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState == nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.VehicleClimater
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector == nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.VehicleClimater
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.Resurrector
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleRange
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleOdometer
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleOdometer
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater == nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleClimater
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleClimater
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleClimater
			api.VehicleRange
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer == nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleClimater
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleClimater
			api.VehicleOdometer
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange == nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleClimater
			api.VehicleOdometer
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
		}

	case chargeState == nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.Resurrector
			api.VehicleClimater
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}

	case chargeState != nil && resurrector != nil && vehicleClimater != nil && vehicleOdometer != nil && vehicleRange != nil:
		return &struct {
			api.Vehicle
			api.ChargeState
			api.Resurrector
			api.VehicleClimater
			api.VehicleOdometer
			api.VehicleRange
		}{
			Vehicle: base,
			ChargeState: &decorateVehicleChargeStateImpl{
				chargeState: chargeState,
			},
			Resurrector: &decorateVehicleResurrectorImpl{
				resurrector: resurrector,
			},
			VehicleClimater: &decorateVehicleVehicleClimaterImpl{
				vehicleClimater: vehicleClimater,
			},
			VehicleOdometer: &decorateVehicleVehicleOdometerImpl{
				vehicleOdometer: vehicleOdometer,
			},
			VehicleRange: &decorateVehicleVehicleRangeImpl{
				vehicleRange: vehicleRange,
			},
		}
	}

	return nil
}

type decorateVehicleChargeStateImpl struct {
	chargeState func() (api.ChargeStatus, error)
}

func (impl *decorateVehicleChargeStateImpl) Status() (api.ChargeStatus, error) {
	return impl.chargeState()
}

type decorateVehicleResurrectorImpl struct {
	resurrector func() error
}

func (impl *decorateVehicleResurrectorImpl) WakeUp() error {
	return impl.resurrector()
}

type decorateVehicleVehicleClimaterImpl struct {
	vehicleClimater func() (bool, error)
}

func (impl *decorateVehicleVehicleClimaterImpl) Climater() (bool, error) {
	return impl.vehicleClimater()
}

type decorateVehicleVehicleOdometerImpl struct {
	vehicleOdometer func() (float64, error)
}

func (impl *decorateVehicleVehicleOdometerImpl) Odometer() (float64, error) {
	return impl.vehicleOdometer()
}

type decorateVehicleVehicleRangeImpl struct {
	vehicleRange func() (int64, error)
}

func (impl *decorateVehicleVehicleRangeImpl) Range() (int64, error) {
	return impl.vehicleRange()
}
