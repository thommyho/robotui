package cmd

import (
	"fmt"
	"strings"

	"github.com/evcc-io/evcc/api"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// vehicleCmd represents the vehicle command
var vehicleCmd = &cobra.Command{
	Use:   "vehicle [name]",
	Short: "Query configured vehicles",
	Run:   runVehicle,
}

func init() {
	rootCmd.AddCommand(vehicleCmd)
	vehicleCmd.PersistentFlags().StringP(flagName, "n", "", fmt.Sprintf(flagNameDescription, "vehicle"))
	vehicleCmd.Flags().BoolP(flagStart, "a", false, flagStartDescription)
	vehicleCmd.Flags().BoolP(flagStop, "o", false, flagStopDescription)
	vehicleCmd.Flags().BoolP(flagWakeup, "w", false, flagWakeupDescription)
	//lint:ignore SA1019 as Title is safe on ascii
	vehicleCmd.Flags().Bool(flagDiagnose, false, strings.Title(flagDiagnose))
}

func runVehicle(cmd *cobra.Command, args []string) {
	// load config
	if err := loadConfigFile(&conf); err != nil {
		fatal(err)
	}

	// setup environment
	if err := configureEnvironment(cmd, conf); err != nil {
		fatal(err)
	}

	// select single vehicle
	if err := selectByName(cmd, &conf.Vehicles); err != nil {
		fatal(err)
	}

	if err := cp.configureVehicles(conf); err != nil {
		fatal(err)
	}

	vehicles := cp.vehicles
	if len(args) == 1 {
		name := args[0]
		vehicle, err := cp.Vehicle(name)
		if err != nil {
			log.FATAL.Fatal(err)
		}

		vehicles = map[string]api.Vehicle{name: vehicle}
	}

	// check single vehicle for error
	if len(vehicles) == 1 {
		if err, ok := maps.Values(vehicles)[0].(error); ok {
			fatal(err)
		}
	}

	var flagUsed bool
	for _, v := range vehicles {
		if cmd.Flags().Lookup(flagWakeup).Changed {
			flagUsed = true

			if vv, ok := v.(api.Resurrector); ok {
				if err := vv.WakeUp(); err != nil {
					log.ERROR.Println("wakeup:", err)
				}
			} else {
				log.ERROR.Println("wakeup: not implemented")
			}
		}

		if cmd.Flags().Lookup(flagStart).Changed {
			flagUsed = true

			if vv, ok := v.(api.VehicleChargeController); ok {
				if err := vv.StartCharge(); err != nil {
					log.ERROR.Println("start charge:", err)
				}
			} else {
				log.ERROR.Println("start charge: not implemented")
			}
		}

		if cmd.Flags().Lookup(flagStop).Changed {
			flagUsed = true

			if vv, ok := v.(api.VehicleChargeController); ok {
				if err := vv.StopCharge(); err != nil {
					log.ERROR.Println("stop charge:", err)
				}
			} else {
				log.ERROR.Println("stop charge: not implemented")
			}
		}
	}

	if !flagUsed {
		d := dumper{len: len(vehicles)}
		flag := cmd.Flags().Lookup(flagDiagnose).Changed

		for name, v := range vehicles {
			d.DumpWithHeader(name, v)
			if flag {
				d.DumpDiagnosis(v)
			}
		}
	}

	// wait for shutdown
	<-shutdownDoneC()
}
