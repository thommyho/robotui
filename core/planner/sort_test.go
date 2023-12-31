package planner

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"github.com/thommyho/robotui/api"
	"golang.org/x/exp/slices"
)

func testRates(clock clock.Clock) api.Rates {
	return api.Rates{
		{
			Price: 2,
			Start: clock.Now().Add(2 * time.Hour),
		},
		{
			Price: 2,
			Start: clock.Now(),
		},
		{
			Price: 1,
			Start: clock.Now().Add(time.Hour),
		},
	}

}
func TestRatesSortByTime(t *testing.T) {
	clock := clock.NewMock()

	r := testRates(clock)

	slices.SortStableFunc(r, SortByTime)
	assert.Equal(t, clock.Now(), r[0].Start)
	assert.Equal(t, clock.Now().Add(time.Hour), r[1].Start) // late slots first
	assert.Equal(t, clock.Now().Add(2*time.Hour), r[2].Start)
}

func TestRatesSortByCost(t *testing.T) {
	clock := clock.NewMock()

	r := testRates(clock)

	slices.SortStableFunc(r, sortByCost)
	assert.Equal(t, clock.Now().Add(time.Hour), r[0].Start)
	assert.Equal(t, clock.Now().Add(2*time.Hour), r[1].Start)
	assert.Equal(t, clock.Now(), r[2].Start)
}
