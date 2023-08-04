package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateCustom(base *Charger, identifier func() (string, error), phaseSwitcher func(int) error, resurrector func() error) api.Charger {
	switch {
	case identifier == nil && phaseSwitcher == nil && resurrector == nil:
		return base

	case identifier != nil && phaseSwitcher == nil && resurrector == nil:
		return &struct {
			*Charger
			api.Identifier
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
		}

	case identifier == nil && phaseSwitcher != nil && resurrector == nil:
		return &struct {
			*Charger
			api.PhaseSwitcher
		}{
			Charger: base,
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier != nil && phaseSwitcher != nil && resurrector == nil:
		return &struct {
			*Charger
			api.Identifier
			api.PhaseSwitcher
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case identifier == nil && phaseSwitcher == nil && resurrector != nil:
		return &struct {
			*Charger
			api.Resurrector
		}{
			Charger: base,
			Resurrector: &decorateCustomResurrectorImpl{
				resurrector: resurrector,
			},
		}

	case identifier != nil && phaseSwitcher == nil && resurrector != nil:
		return &struct {
			*Charger
			api.Identifier
			api.Resurrector
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
			Resurrector: &decorateCustomResurrectorImpl{
				resurrector: resurrector,
			},
		}

	case identifier == nil && phaseSwitcher != nil && resurrector != nil:
		return &struct {
			*Charger
			api.PhaseSwitcher
			api.Resurrector
		}{
			Charger: base,
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
			Resurrector: &decorateCustomResurrectorImpl{
				resurrector: resurrector,
			},
		}

	case identifier != nil && phaseSwitcher != nil && resurrector != nil:
		return &struct {
			*Charger
			api.Identifier
			api.PhaseSwitcher
			api.Resurrector
		}{
			Charger: base,
			Identifier: &decorateCustomIdentifierImpl{
				identifier: identifier,
			},
			PhaseSwitcher: &decorateCustomPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
			Resurrector: &decorateCustomResurrectorImpl{
				resurrector: resurrector,
			},
		}
	}

	return nil
}

type decorateCustomIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateCustomIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateCustomPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateCustomPhaseSwitcherImpl) Phases1p3p(phases int) error {
	return impl.phaseSwitcher(phases)
}

type decorateCustomResurrectorImpl struct {
	resurrector func() error
}

func (impl *decorateCustomResurrectorImpl) WakeUp() error {
	return impl.resurrector()
}