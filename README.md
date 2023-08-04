# robotui 🚘☀️

[![Build Status](https://github.com/thommyho/robotui/workflows/Build/badge.svg)](https://github.com/thommyho/robotui/actions?query=workflow%3ABuild)
[![Code Quality](https://goreportcard.com/badge/github.com/thommyho/robotui)](https://goreportcard.com/report/github.com/thommyho/robotui)
[![Translation](https://hosted.weblate.org/widgets/robotui/-/robotui/svg-badge.svg)](https://hosted.weblate.org/engage/robotui/)
[![Open in Visual Studio Code](https://img.shields.io/static/v1?logo=visualstudiocode&label=&message=Open%20in%20VS%20Code&labelColor=2c2c32&color=007acc&logoColor=007acc)](https://open.vscode.dev/evcc-io/robotui)
[![OSS hosting by cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith)](https://cloudsmith.io/~robotui/packages/)
[![Latest Version](https://img.shields.io/github/release/evcc-io/robotui.svg)](https://github.com/thommyho/robotui/releases)

robotui is an extensible EV Charge Controller with PV integration implemented in [Go][1]. Featured in [PV magazine](https://www.pv-magazine.de/2021/01/15/selbst-ist-der-groeoenlandhof-wallbox-ladesteuerung-selbst-gebaut/).

![Screenshot](docs/screenshot.png)

## Features

- simple and clean user interface
- wide range of supported [chargers](https://docs.robotui.io/docs/devices/chargers):
  - ABL eMH1, Alfen (Eve), Bender (CC612/613), cFos (PowerBrain), Daheimladen, Ebee (Wallbox), Ensto (Chago Wallbox), [EVSEWifi/ smartWB](https://www.evse-wifi.de), Garo (GLB, GLB+, LS4), go-eCharger, HardyBarth (eCB1, cPH1, cPH2), Heidelberg (Energy Control), Innogy (eBox), Juice (Charger Me), KEBA/BMW, Menneckes (Amedio, Amtron Premium/Xtra, Amtron ChargeConrol), NRGkick, [openWB (includes Pro)](https://openwb.de/), Optec (Mobility One), PC Electric (includes Garo), Siemens, TechniSat (Technivolt), [Tinkerforge Warp Charger](https://www.warp-charger.com), Ubitricity (Heinz), Vestel, Wallbe, Webasto (Live), Mobile Charger Connect and many more
  - experimental EEBus support (Elli, PMCC)
  - experimental OCPP support
  - Build-your-own: Phoenix Contact (includes ESL Walli), [EVSE DIN](http://evracing.cz/simple-evse-wallbox)
  - Smart-Home outlets: FritzDECT, Shelly, Tasmota, TP-Link
- wide range of supported [meters](https://docs.robotui.io/docs/devices/meters) for grid, pv, battery and charger:
  - ModBus: Eastron SDM, MPM3PM, ORNO WE, SBC ALE3 and many more, see <https://github.com/volkszaehler/mbmd#supported-devices> for a complete list
  - Integrated systems: SMA Sunny Home Manager and Energy Meter, KOSTAL Smart Energy Meter (KSEM, EMxx)
  - Sunspec-compatible inverter or home battery devices: Fronius, SMA, SolarEdge, KOSTAL, STECA, E3DC, ...
  - and various others: Discovergy, Tesla PowerWall, LG ESS HOME, OpenEMS (FENECON)
- [vehicle](https://docs.robotui.io/docs/devices/vehicles) integration (state of charge, remote charge, battery and preconditioning status):
  - Audi, BMW, Citroën, Dacia, Fiat, Ford, Hyundai, Jaguar, Kia, Landrover, Mercedes, Mini, Nissan, Opel, Peugeot, Porsche, Renault, Seat, Smart, Skoda, Tesla, Volkswagen, Volvo, ...
  - Services: OVMS, Tronity
  - Scooters: Niu, Silence
- [plugins](https://docs.robotui.io/docs/reference/plugins) for integrating with any charger/ meter/ vehicle:
  - Modbus, HTTP, MQTT, Javascript, WebSockets and shell scripts
- status [notifications](https://docs.robotui.io/docs/reference/configuration/messaging) using [Telegram](https://telegram.org), [PushOver](https://pushover.net) and [many more](https://containrrr.dev/shoutrrr/)
- logging using [InfluxDB](https://www.influxdata.com) and [Grafana](https://grafana.com/grafana/)
- granular charge power control down to mA steps with supported chargers (labeled by e.g. smartWB as [OLC](https://board.evse-wifi.de/viewtopic.php?f=16&t=187))
- REST and MQTT [APIs](https://docs.robotui.io/docs/reference/api) for integration with home automation systems
- Add-ons for [HomeAssistant](https://github.com/thommyho/robotui-hassio-addon) and [OpenHAB](https://www.openhab.org/addons/bindings/robotui) (not maintained by the robotui core team)

## Getting Started

You'll find everything you need in our [documentation](https://docs.robotui.io/) (German).

## Contribute

To build robotui from source, [Go][1] 1.20 and [Node][2] 18 are required.

Build and run go backend. The UI becomes available at http://127.0.0.1:7070/

```sh
make install-ui
make ui
make install
make
./robotui
```

### UI development

For frontend development start the Vue toolchain in dev-mode. Open http://127.0.0.1:7071/ to get to the livelreloading development server. It pulls its data from port 7070 (see above).

```sh
npm install
npm run start
```

### Integration tests

We use Playwright for end-to-end integration tests. They start a local robotui instance with different configuration yamls and prefilled databases. To run them, you have to do a local build first.

```sh
make ui build
npm run playwright
```

### Code formatting

We use linters (golangci-lint, Prettier) to keep a coherent source code formatting. It's recommended to use the format-on-save feature of your editor. For VSCode use the [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go), [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode) and [Veture](https://marketplace.visualstudio.com/items?itemName=octref.vetur) extension. You can manually reformat your code by running:

```sh
make lint
make lint-ui
```

### Changing templates

robotui supports a massive amount of different devices. To keep our documentation and website in sync with the latest software the core project (this repo) generates meta-data that's pushed to the `docs` and `robotui.io` repository. Make sure to update this meta-data every time you make changes to a templates.

```sh
make docs
```

If you miss one of the above steps Gitub Actions will likely trigger a **Porcelain** error.

### Adding or modifying translations

robotui already includes many translations for the UI. Weblate Hosted is used to maintain all languages. Feel free to add more languages or verify and edit existing translations. Weblate will automatically push all modifications on a regular base to the robotui repository.

[![Weblate Hosted](https://hosted.weblate.org/widgets/robotui/-/robotui/287x66-grey.png)](https://hosted.weblate.org/engage/robotui/)
[![Languages](https://hosted.weblate.org/widgets/robotui/-/robotui/multi-auto.svg)](https://hosted.weblate.org/engage/robotui/)

https://hosted.weblate.org/projects/robotui/robotui/

## Sponsorship

<img src="docs/logo.png" align="right" width="150" />

robotui believes in open source software. We're committed to provide best in class EV charging experience.
Maintaining robotui consumes time and effort. With the vast amount of different devices to support, we depend on community and vendor support to keep robotui alive.

While robotui is open source, we would also like to encourage vendors to provide open source hardware devices, public documentation and support open source projects like ours that provide additional value to otherwise closed hardware. Where this is not the case, robotui requires "sponsor token" to finance ongoing development and support of robotui.

The personal sponsor token requires a [Github Sponsorship](https://github.com/sponsors/evcc-io) and can be requested at [sponsor.robotui.io](https://sponsor.robotui.io/).

[1]: https://golang.org
[2]: https://nodejs.org/
