package tariff

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/tariff/elering"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/request"
	"golang.org/x/exp/slices"
)

type Elering struct {
	*embed
	mux     sync.Mutex
	log     *util.Logger
	region  string
	data    api.Rates
	updated time.Time
}

var _ api.Tariff = (*Elering)(nil)

func init() {
	registry.Add("elering", NewEleringFromConfig)
}

func NewEleringFromConfig(other map[string]interface{}) (api.Tariff, error) {
	var cc struct {
		embed  `mapstructure:",squash"`
		Region string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Region == "" {
		return nil, errors.New("missing region")
	}

	t := &Elering{
		embed:  &cc.embed,
		log:    util.NewLogger("Elering"),
		region: strings.ToLower(cc.Region),
	}

	done := make(chan error)
	go t.run(done)
	err := <-done

	return t, err
}

func (t *Elering) run(done chan error) {
	var once sync.Once
	client := request.NewHelper(t.log)
	bo := newBackoff()

	for ; true; <-time.Tick(time.Hour) {
		var res elering.NpsPrice

		ts := time.Now().Truncate(time.Hour)
		uri := fmt.Sprintf("%s/nps/price?start=%s&end=%s", elering.URI,
			url.QueryEscape(ts.Format(time.RFC3339)),
			url.QueryEscape(ts.Add(48*time.Hour).Format(time.RFC3339)))

		if err := backoff.Retry(func() error {
			return client.GetJSON(uri, &res)
		}, bo); err != nil {
			once.Do(func() { done <- err })

			t.log.ERROR.Println(err)
			continue
		}

		once.Do(func() { close(done) })

		t.mux.Lock()
		t.updated = time.Now()

		data := res.Data[t.region]

		t.data = make(api.Rates, 0, len(data))
		for _, r := range data {
			ts := time.Unix(r.Timestamp, 0)

			ar := api.Rate{
				Start: ts.Local(),
				End:   ts.Add(time.Hour).Local(),
				Price: t.totalPrice(r.Price / 1e3),
			}
			t.data = append(t.data, ar)
		}

		t.mux.Unlock()
	}
}

// Rates implements the api.Tariff interface
func (t *Elering) Rates() (api.Rates, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	return slices.Clone(t.data), outdatedError(t.updated, time.Hour)
}

// Type implements the api.Tariff interface
func (t *Elering) Type() api.TariffType {
	return api.TariffTypePriceDynamic
}
