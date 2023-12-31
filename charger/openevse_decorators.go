package charger

// Code generated by github.com/thommyho/robotui/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/thommyho/robotui/api"
)

func decorateOpenEVSE(base *OpenEVSE, phaseSwitcher func(phases int) error) api.Charger {
	switch {
	case phaseSwitcher == nil:
		return base

	case phaseSwitcher != nil:
		return &struct {
			*OpenEVSE
			api.PhaseSwitcher
		}{
			OpenEVSE: base,
			PhaseSwitcher: &decorateOpenEVSEPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateOpenEVSEPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateOpenEVSEPhaseSwitcherImpl) Phases1p3p(phases int) error {
	return impl.phaseSwitcher(phases)
}
