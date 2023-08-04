package core

import (
	"github.com/thommyho/robotui/api"
	"golang.org/x/exp/slices"
)

// chargerHasFeature checks availability of charger feature
func (lp *Loadpoint) chargerHasFeature(f api.Feature) bool {
	c, ok := lp.charger.(api.FeatureDescriber)
	if ok {
		ok = slices.Contains(c.Features(), f)
	}
	return ok
}

// publishChargerFeature publishes availability of charger features
func (lp *Loadpoint) publishChargerFeature(f api.Feature) {
	c, ok := lp.charger.(api.FeatureDescriber)
	if ok {
		ok = slices.Contains(c.Features(), f)
	}
	lp.publish("chargerFeature"+f.String(), ok)
}

// chargerSoc returns charger soc if available
func (lp *Loadpoint) chargerSoc() (float64, error) {
	if c, ok := lp.charger.(api.Battery); ok {
		return c.Soc()
	}
	return 0, api.ErrNotAvailable
}
