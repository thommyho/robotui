package meter

import (
	"errors"
	"fmt"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/provider"
	"github.com/evcc-io/evcc/util"
)

func init() {
	registry.Add(api.Custom, NewConfigurableFromConfig)
}

//go:generate go run ../cmd/tools/decorate.go -f decorateMeter -b api.Meter -t "api.MeterEnergy,TotalEnergy,func() (float64, error)" -t "api.PhaseCurrents,Currents,func() (float64, float64, float64, error)" -t "api.PhaseVoltages,Voltages,func() (float64, float64, float64, error)" -t "api.PhasePowers,Powers,func() (float64, float64, float64, error)" -t "api.Battery,Soc,func() (float64, error)" -t "api.BatteryCapacity,Capacity,func() float64"

// NewConfigurableFromConfig creates api.Meter from config
func NewConfigurableFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		capacity `mapstructure:",squash"`
		Power    provider.Config
		Energy   *provider.Config  // optional
		Soc      *provider.Config  // optional
		Currents []provider.Config // optional
		Voltages []provider.Config // optional
		Powers   []provider.Config // optional
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	power, err := provider.NewFloatGetterFromConfig(cc.Power)
	if err != nil {
		return nil, fmt.Errorf("power: %w", err)
	}

	m, _ := NewConfigurable(power)

	// decorate energy
	var totalEnergyG func() (float64, error)
	if cc.Energy != nil {
		totalEnergyG, err = provider.NewFloatGetterFromConfig(*cc.Energy)
		if err != nil {
			return nil, fmt.Errorf("energy: %w", err)
		}
	}

	// decorate currents
	currentsG, err := buildPhaseProviders(cc.Currents)
	if err != nil {
		return nil, fmt.Errorf("currents: %w", err)
	}

	// decorate voltages
	voltagesG, err := buildPhaseProviders(cc.Voltages)
	if err != nil {
		return nil, fmt.Errorf("voltages: %w", err)
	}

	// decorate powers
	powersG, err := buildPhaseProviders(cc.Powers)
	if err != nil {
		return nil, fmt.Errorf("powers: %w", err)
	}

	// decorate soc
	var batterySocG func() (float64, error)
	if cc.Soc != nil {
		batterySocG, err = provider.NewFloatGetterFromConfig(*cc.Soc)
		if err != nil {
			return nil, fmt.Errorf("battery: %w", err)
		}
	}

	res := m.Decorate(totalEnergyG, currentsG, voltagesG, powersG, batterySocG, cc.capacity.Decorator())

	return res, nil
}

func buildPhaseProviders(providers []provider.Config) (func() (float64, float64, float64, error), error) {
	var res func() (float64, float64, float64, error)
	if len(providers) > 0 {
		if len(providers) != 3 {
			return nil, errors.New("need one per phase, total three")
		}

		phases := make([]func() (float64, error), 0, 3)
		for idx, prov := range providers {
			c, err := provider.NewFloatGetterFromConfig(prov)
			if err != nil {
				return nil, fmt.Errorf("[%d] %w", idx, err)
			}

			phases = append(phases, c)
		}

		res = collectPhaseProviders(phases)
	}

	return res, nil
}

// collectPhaseProviders combines phase getters into currents api function
func collectPhaseProviders(g []func() (float64, error)) func() (float64, float64, float64, error) {
	return func() (float64, float64, float64, error) {
		var res []float64
		for _, currentG := range g {
			c, err := currentG()
			if err != nil {
				return 0, 0, 0, err
			}

			res = append(res, c)
		}

		return res[0], res[1], res[2], nil
	}
}

// NewConfigurable creates a new meter
func NewConfigurable(currentPowerG func() (float64, error)) (*Meter, error) {
	m := &Meter{
		currentPowerG: currentPowerG,
	}
	return m, nil
}

// Meter is an api.Meter implementation with configurable getters and setters.
type Meter struct {
	currentPowerG func() (float64, error)
}

// Decorate attaches additional capabilities to the base meter
func (m *Meter) Decorate(
	totalEnergy func() (float64, error),
	currents func() (float64, float64, float64, error),
	voltages func() (float64, float64, float64, error),
	powers func() (float64, float64, float64, error),
	batterySoc func() (float64, error),
	capacity func() float64,
) api.Meter {
	return decorateMeter(m, totalEnergy, currents, voltages, powers, batterySoc, capacity)
}

// CurrentPower implements the api.Meter interface
func (m *Meter) CurrentPower() (float64, error) {
	return m.currentPowerG()
}
