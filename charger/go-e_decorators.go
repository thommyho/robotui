package charger

// Code generated by github.com/andig/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/thommyho/robotui/api"
)

func decorateGoE(base *GoE, phaseSwitcher func(phases int) error) api.Charger {
	switch {
	case phaseSwitcher == nil:
		return base

	case phaseSwitcher == nil:
		return &struct {
			*GoE
			api.MeterEnergy
		}{
			GoE: base,
		}

	case phaseSwitcher != nil:
		return &struct {
			*GoE
			api.PhaseSwitcher
		}{
			GoE: base,
			PhaseSwitcher: &decorateGoEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case phaseSwitcher != nil:
		return &struct {
			*GoE
			api.PhaseSwitcher
			api.MeterEnergy
		}{
			GoE: base,
			PhaseSwitcher: &decorateGoEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateGoEPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateGoEPhaseSwitcherImpl) Phases1p3p(phases int) error {
	return impl.phaseSwitcher(phases)
}
