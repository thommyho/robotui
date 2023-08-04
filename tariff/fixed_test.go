package tariff

import (
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/golang-module/carbon/v2"
	"github.com/robotuimyhorobotuiotui/tariff/fixed"
	"github.com/stretchr/testify/assert"
	"github.com/thommyho/robotui/api"
)

func TestFixed(t *testing.T) {
	tf := &Fixed{
		clock: clock.NewMock(),
		zones: []fixed.Zone{
			{Price: 0.3},
		},
	}

	var expect api.Rates
	for dow := 0; dow < 7; dow++ {
		dayStart := carbon.FromStdTime(tf.clock.Now()).StartOfDay().AddDays(dow)

		for hour := 0; hour < 24; hour++ {
			expect = append(expect, api.Rate{
				Price: 0.3,
				Start: dayStart.AddHours(hour).ToStdTime(),
				End:   dayStart.AddHours(hour + 1).ToStdTime(),
			})
		}
	}

	rates, err := tf.Rates()
	assert.NoError(t, err)
	assert.Equal(t, expect, rates)
}

func TestFixedSplitZones(t *testing.T) {
	at, err := NewFixedFromConfig(map[string]interface{}{
		"price": 0.5,
		"zones": []struct {
			Price float64
			Hours string
		}{
			{0.1, "0-5:30,21-0"},
		},
	})
	assert.NoError(t, err)

	tf := at.(*Fixed)
	tf.clock = clock.NewMock()

	var expect api.Rates
	for i := 0; i < 7; i++ {
		dayStart := carbon.FromStdTime(tf.clock.Now()).StartOfDay().AddDays(i)

		// 00:00-05:00 0.1
		for hour := 0; hour < 5; hour++ {
			expect = append(expect, api.Rate{
				Price: 0.1,
				Start: dayStart.AddHours(hour).ToStdTime(),
				End:   dayStart.AddHours(hour + 1).ToStdTime(),
			})
		}

		// 05:00-05:30 0.1
		expect = append(expect, api.Rate{
			Price: 0.1,
			Start: dayStart.AddHours(5).ToStdTime(),
			End:   dayStart.AddHours(5).AddMinutes(30).ToStdTime(),
		})

		// 05:30-06:00 0.5
		expect = append(expect, api.Rate{
			Price: 0.5,
			Start: dayStart.AddHours(5).AddMinutes(30).ToStdTime(),
			End:   dayStart.AddHours(6).ToStdTime(),
		})

		// 06:00-21:00 0.5
		for hour := 6; hour < 21; hour++ {
			expect = append(expect, api.Rate{
				Price: 0.5,
				Start: dayStart.AddHours(hour).ToStdTime(),
				End:   dayStart.AddHours(hour + 1).ToStdTime(),
			})
		}

		// 21:00-00:00 0.1
		for hour := 21; hour < 24; hour++ {
			expect = append(expect, api.Rate{
				Price: 0.1,
				Start: dayStart.AddHours(hour).ToStdTime(),
				End:   dayStart.AddHours(hour + 1).ToStdTime(),
			})
		}
	}

	rates, err := tf.Rates()
	assert.NoError(t, err)
	assert.Equal(t, expect, rates)
}
