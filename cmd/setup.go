package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/libp2p/zeroconf/v2"
	"github.com/robotuimyhorobotuiotui/charger/eebus"
	"github.com/robotuimyhorobotuiotui/cmd/shutdown"
	"github.com/robotuimyhorobotuiotui/core"
	"github.com/robotuimyhorobotuiotui/core/site"
	"github.com/robotuimyhorobotuiotui/hems"
	"github.com/robotuimyhorobotuiotui/provider/golang"
	"github.com/robotuimyhorobotuiotui/provider/javascript"
	"github.com/robotuimyhorobotuiotui/provider/mqtt"
	"github.com/robotuimyhorobotuiotui/push"
	"github.com/robotuimyhorobotuiotui/server"
	"github.com/robotuimyhorobotuiotui/server/db"
	"github.com/robotuimyhorobotuiotui/server/db/settings"
	"github.com/robotuimyhorobotuiotui/tariff"
	"github.com/robotuimyhorobotuiotui/util"
	"github.com/robotuimyhorobotuiotui/util/locale"
	"github.com/robotuimyhorobotuiotui/util/machine"
	"github.com/robotuimyhorobotuiotui/util/pipe"
	"github.com/robotuimyhorobotuiotui/util/request"
	"github.com/robotuimyhorobotuiotui/util/sponsor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thommyho/robotui/api"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/text/currency"
)

var cp = new(ConfigProvider)

func loadConfigFile(conf *config) error {
	err := viper.ReadInConfig()

	if cfgFile = viper.ConfigFileUsed(); cfgFile == "" {
		return err
	}

	log.INFO.Println("using config file:", cfgFile)

	if err == nil {
		if err = viper.UnmarshalExact(&conf); err != nil {
			err = fmt.Errorf("failed parsing config file: %w", err)
		}
	}

	// parse log levels after reading config
	if err == nil {
		parseLogLevels()
	}

	return err
}

func configureEnvironment(cmd *cobra.Command, conf config) (err error) {
	// full http request log
	if cmd.Flags().Lookup(flagHeaders).Changed {
		request.LogHeaders = true
	}

	// setup machine id
	if conf.Plant != "" {
		err = machine.CustomID(conf.Plant)
	}

	// setup sponsorship (allow env override)
	if err == nil && conf.SponsorToken != "" {
		err = sponsor.ConfigureSponsorship(conf.SponsorToken)
	}

	// setup translations
	if err == nil {
		err = locale.Init()
	}

	// setup persistence
	if err == nil && conf.Database.Dsn != "" {
		err = configureDatabase(conf.Database)
	}

	// setup mqtt client listener
	if err == nil && conf.Mqtt.Broker != "" {
		err = configureMQTT(conf.Mqtt)
	}

	// setup javascript VMs
	if err == nil {
		err = configureJavascript(conf.Javascript)
	}

	// setup go VMs
	if err == nil {
		err = configureGo(conf.Go)
	}

	// setup EEBus server
	if err == nil && conf.EEBus != nil {
		err = configureEEBus(conf.EEBus)
	}

	return
}

// configureDatabase configures session database
func configureDatabase(conf dbConfig) error {
	if err := db.NewInstance(conf.Type, conf.Dsn); err != nil {
		return err
	}

	if err := settings.Init(); err != nil {
		return err
	}

	shutdown.Register(func() {
		if err := settings.Persist(); err != nil {
			log.ERROR.Println("cannot save settings:", err)
		}
	})

	return nil
}

// configureInflux configures influx database
func configureInflux(conf server.InfluxConfig, site site.API, in <-chan util.Param) {
	influx := server.NewInfluxClient(
		conf.URL,
		conf.Token,
		conf.Org,
		conf.User,
		conf.Password,
		conf.Database,
	)

	// eliminate duplicate values
	dedupe := pipe.NewDeduplicator(30*time.Minute, "vehicleCapacity", "vehicleSoc", "vehicleRange", "vehicleOdometer", "chargedEnergy", "chargeRemainingEnergy")
	in = dedupe.Pipe(in)

	go influx.Run(site, in)
}

// setup mqtt
func configureMQTT(conf mqttConfig) error {
	log := util.NewLogger("mqtt")

	var err error
	if mqtt.Instance, err = mqtt.RegisteredClient(log, conf.Broker, conf.User, conf.Password, conf.ClientID, 1, conf.Insecure, func(options *paho.ClientOptions) {
		topic := fmt.Sprintf("%s/status", strings.Trim(conf.Topic, "/"))
		options.SetWill(topic, "offline", 1, true)
	}); err != nil {
		return fmt.Errorf("failed configuring mqtt: %w", err)
	}

	return nil
}

// setup javascript
func configureJavascript(conf []javascriptConfig) error {
	for _, cc := range conf {
		if _, err := javascript.RegisteredVM(cc.VM, cc.Script); err != nil {
			return fmt.Errorf("failed configuring javascript: %w", err)
		}
	}
	return nil
}

// setup go
func configureGo(conf []goConfig) error {
	for _, cc := range conf {
		if _, err := golang.RegisteredVM(cc.VM, cc.Script); err != nil {
			return fmt.Errorf("failed configuring go: %w", err)
		}
	}
	return nil
}

// setup HEMS
func configureHEMS(conf typedConfig, site *core.Site, httpd *server.HTTPd) error {
	hems, err := hems.NewFromConfig(conf.Type, conf.Other, site, httpd)
	if err != nil {
		return fmt.Errorf("failed configuring hems: %w", err)
	}

	go hems.Run()

	return nil
}

// setup MDNS
func configureMDNS(conf networkConfig) error {
	host := strings.TrimSuffix(conf.Host, ".local")

	zc, err := zeroconf.RegisterProxy("EV Charge Controller", "_http._tcp", "local.", conf.Port, host, nil, []string{}, nil)
	if err != nil {
		return fmt.Errorf("mDNS announcement: %w", err)
	}

	shutdown.Register(zc.Shutdown)

	return nil
}

// setup EEBus
func configureEEBus(conf map[string]interface{}) error {
	var err error
	if eebus.Instance, err = eebus.NewServer(conf); err != nil {
		return fmt.Errorf("failed configuring eebus: %w", err)
	}

	eebus.Instance.Run()
	shutdown.Register(eebus.Instance.Shutdown)

	return nil
}

// setup messaging
func configureMessengers(conf messagingConfig, valueChan chan util.Param, cache *util.Cache) (chan push.Event, error) {
	messageChan := make(chan push.Event, 1)

	messageHub, err := push.NewHub(conf.Events, cache)
	if err != nil {
		return messageChan, fmt.Errorf("failed configuring push services: %w", err)
	}

	for _, service := range conf.Services {
		impl, err := push.NewFromConfig(service.Type, service.Other)
		if err != nil {
			return messageChan, fmt.Errorf("failed configuring push service %s: %w", service.Type, err)
		}
		messageHub.Add(impl)
	}

	go messageHub.Run(messageChan, valueChan)

	return messageChan, nil
}

func configureTariffs(conf tariffConfig) (tariff.Tariffs, error) {
	var grid, feedin, co2, planner api.Tariff
	var currencyCode currency.Unit = currency.EUR
	var err error

	if conf.Currency != "" {
		currencyCode = currency.MustParseISO(conf.Currency)
	}

	if conf.Grid.Type != "" {
		grid, err = tariff.NewFromConfig(conf.Grid.Type, conf.Grid.Other)
		if err != nil {
			grid = nil
			log.ERROR.Printf("failed configuring grid tariff: %v", err)
		}
	}

	if conf.FeedIn.Type != "" {
		feedin, err = tariff.NewFromConfig(conf.FeedIn.Type, conf.FeedIn.Other)
		if err != nil {
			feedin = nil
			log.ERROR.Printf("failed configuring feed-in tariff: %v", err)
		}
	}

	if conf.Co2.Type != "" {
		co2, err = tariff.NewFromConfig(conf.Co2.Type, conf.Co2.Other)
		if err != nil {
			co2 = nil
			log.ERROR.Printf("failed configuring co2 tariff: %v", err)
		}
	}

	if conf.Planner.Type != "" {
		planner, err = tariff.NewFromConfig(conf.Planner.Type, conf.Planner.Other)
		if err != nil {
			planner = nil
			log.ERROR.Printf("failed configuring planner tariff: %v", err)
		} else if planner.Type() == api.TariffTypeCo2 {
			log.WARN.Printf("tariff configuration changed, use co2 instead of planner for co2 tariff")
		}
	}

	tariffs := tariff.NewTariffs(currencyCode, grid, feedin, co2, planner)

	return *tariffs, nil
}

func configureSiteAndLoadpoints(conf config) (*core.Site, error) {
	if err := cp.configure(conf); err != nil {
		return nil, err
	}

	loadpoints, err := configureLoadpoints(conf, cp)
	if err != nil {
		return nil, fmt.Errorf("failed configuring loadpoints: %w", err)
	}

	tariffs, err := configureTariffs(conf.Tariffs)
	if err != nil {
		return nil, err
	}

	// list of vehicles ordered by name
	keys := maps.Keys(cp.vehicles)
	slices.Sort(keys)

	vehicles := make([]api.Vehicle, 0, len(cp.vehicles))
	for _, k := range keys {
		vehicles = append(vehicles, cp.vehicles[k])
	}

	return configureSite(conf.Site, cp, loadpoints, vehicles, tariffs)
}

func configureSite(conf map[string]interface{}, cp *ConfigProvider, loadpoints []*core.Loadpoint, vehicles []api.Vehicle, tariffs tariff.Tariffs) (*core.Site, error) {
	site, err := core.NewSiteFromConfig(log, cp, conf, loadpoints, vehicles, tariffs)
	if err != nil {
		return nil, fmt.Errorf("failed configuring site: %w", err)
	}

	return site, nil
}

func configureLoadpoints(conf config, cp *ConfigProvider) (loadpoints []*core.Loadpoint, err error) {
	lpInterfaces, ok := viper.AllSettings()["loadpoints"].([]interface{})
	if !ok || len(lpInterfaces) == 0 {
		return nil, errors.New("missing loadpoints")
	}

	for id, lpcI := range lpInterfaces {
		var lpc map[string]interface{}
		if err := util.DecodeOther(lpcI, &lpc); err != nil {
			return nil, fmt.Errorf("failed decoding loadpoint configuration: %w", err)
		}

		log := util.NewLogger("lp-" + strconv.Itoa(id+1))
		lp, err := core.NewLoadpointFromConfig(log, cp, lpc)
		if err != nil {
			return nil, fmt.Errorf("failed configuring loadpoint: %w", err)
		}

		loadpoints = append(loadpoints, lp)
	}

	return loadpoints, nil
}
