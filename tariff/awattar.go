package tariff

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/tariff/awattar"
	"github.com/thommyho/robotui/util"
	"github.com/thommyho/robotui/util/request"
	"golang.org/x/exp/slices"
)

type Awattar struct {
	*embed
	mux     sync.Mutex
	log     *util.Logger
	uri     string
	data    api.Rates
	updated time.Time
}

var _ api.Tariff = (*Awattar)(nil)

func init() {
	registry.Add("awattar", NewAwattarFromConfig)
}

func NewAwattarFromConfig(other map[string]interface{}) (api.Tariff, error) {
	cc := struct {
		embed  `mapstructure:",squash"`
		Region string
	}{
		Region: "DE",
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	t := &Awattar{
		embed: &cc.embed,
		log:   util.NewLogger("awattar"),
		uri:   fmt.Sprintf(awattar.RegionURI, strings.ToLower(cc.Region)),
	}

	done := make(chan error)
	go t.run(done)
	err := <-done

	return t, err
}

func (t *Awattar) run(done chan error) {
	var once sync.Once
	bo := newBackoff()
	client := request.NewHelper(t.log)

	for ; true; <-time.Tick(time.Hour) {
		var res awattar.Prices

		if err := backoff.Retry(func() error {
			return client.GetJSON(t.uri, &res)
		}, bo); err != nil {
			once.Do(func() { done <- err })

			t.log.ERROR.Println(err)
			continue
		}

		once.Do(func() { close(done) })

		t.mux.Lock()
		t.updated = time.Now()

		t.data = make(api.Rates, 0, len(res.Data))
		for _, r := range res.Data {
			ar := api.Rate{
				Start: r.StartTimestamp.Local(),
				End:   r.EndTimestamp.Local(),
				Price: t.totalPrice(r.Marketprice / 1e3),
			}
			t.data = append(t.data, ar)
		}

		t.mux.Unlock()
	}
}

// Rates implements the api.Tariff interface
func (t *Awattar) Rates() (api.Rates, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	return slices.Clone(t.data), outdatedError(t.updated, time.Hour)
}

// Type implements the api.Tariff interface
func (t *Awattar) Type() api.TariffType {
	return api.TariffTypePriceDynamic
}
