package tariff

import (
	"time"

	"github.com/thommyho/robotui/api"
	"golang.org/x/text/currency"
)

type Tariffs struct {
	Currency                   currency.Unit
	Grid, FeedIn, Co2, Planner api.Tariff
}

func NewTariffs(currency currency.Unit, grid, feedin, co2 api.Tariff, planner api.Tariff) *Tariffs {
	return &Tariffs{
		Currency: currency,
		Grid:     grid,
		FeedIn:   feedin,
		Co2:      co2,
		Planner:  planner,
	}
}

func currentPrice(t api.Tariff) (float64, error) {
	if t != nil {
		if rr, err := t.Rates(); err == nil {
			if r, err := rr.Current(time.Now()); err == nil {
				return r.Price, nil
			}
		}
	}
	return 0, api.ErrNotAvailable
}

// CurrentGridPrice returns the current grid price.
func (t *Tariffs) CurrentGridPrice() (float64, error) {
	return currentPrice(t.Grid)
}

// CurrentFeedInPrice returns the current feed-in price.
func (t *Tariffs) CurrentFeedInPrice() (float64, error) {
	return currentPrice(t.FeedIn)
}

// CurrentCo2 determines the grids co2 emission.
func (t *Tariffs) CurrentCo2() (float64, error) {
	if t.Co2 != nil {
		return currentPrice(t.Co2)
	}
	return 0, api.ErrNotAvailable
}

// outdatedError returns api.ErrOutdated if t is older than 2*d
func outdatedError(t time.Time, d time.Duration) error {
	if time.Since(t) > 2*d {
		return api.ErrOutdated
	}
	return nil
}
