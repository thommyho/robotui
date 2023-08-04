package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/thommyho/robotui/cmd/shutdown"
	"github.com/thommyho/robotui/util"
)

// parseLogLevels parses --log area:level[,...] switch into levels per log area
func parseLogLevels() {
	levels := viper.GetStringMapString("levels")

	var level string
	for _, kv := range strings.Split(viper.GetString("log"), ",") {
		areaLevel := strings.SplitN(kv, ":", 2)
		if len(areaLevel) == 1 {
			level = areaLevel[0]
		} else {
			levels[areaLevel[0]] = areaLevel[1]
		}
	}

	util.LogLevel(level, levels)
}

// unwrap converts a wrapped error into slice of strings
func unwrap(err error) (res []string) {
	for err != nil {
		inner := errors.Unwrap(err)
		if inner == nil {
			res = append(res, err.Error())
		} else {
			cur := strings.TrimSuffix(err.Error(), ": "+inner.Error())
			cur = strings.TrimSuffix(cur, inner.Error())
			res = append(res, strings.TrimSpace(cur))
		}
		err = inner
	}
	return
}

// redact redacts a configuration string
func redact(src string) string {
	secrets := []string{
		"mac",                   // infrastructure
		"sponsortoken", "plant", // global settings
		"user", "password", "pin", // users
		"token", "access", "refresh", "accesstoken", "refreshtoken", // tokens, including template variations
		"ain", "secret", "serial", "deviceid", "machineid", "idtag", // devices
		"app", "chats", "recipients", // push messaging
		"vin"} // vehicles
	return regexp.
		MustCompile(fmt.Sprintf(`(?i)\b(%s)\b.*?:.*`, strings.Join(secrets, "|"))).
		ReplaceAllString(src, "$1: *****")
}

func publishErrorInfo(valueChan chan<- util.Param, cfgFile string, err error) {
	if cfgFile != "" {
		file, pathErr := filepath.Abs(cfgFile)
		if pathErr != nil {
			file = cfgFile
		}
		valueChan <- util.Param{Key: "file", Val: file}

		if src, fileErr := os.ReadFile(cfgFile); fileErr != nil {
			log.ERROR.Println("could not open config file:", fileErr)
		} else {
			valueChan <- util.Param{Key: "config", Val: redact(string(src))}

			// find line number
			if match := regexp.MustCompile(`yaml: line (\d+):`).FindStringSubmatch(err.Error()); len(match) == 2 {
				if line, err := strconv.Atoi(match[1]); err == nil {
					valueChan <- util.Param{Key: "line", Val: line}
				}
			}
		}
	}

	valueChan <- util.Param{Key: "fatal", Val: unwrap(err)}
}

// fatal logs a fatal error and runs shutdown functions before terminating
func fatal(err error) {
	log.FATAL.Println(err)
	<-shutdownDoneC()
	os.Exit(1)
}

// shutdownDoneC returns a channel that closes when shutdown has completed
func shutdownDoneC() <-chan struct{} {
	doneC := make(chan struct{})
	go shutdown.Cleanup(doneC)
	return doneC
}

func wrapErrors(err error) error {
	if err != nil {
		var opErr *net.OpError
		var pathErr *os.PathError

		switch {
		case errors.As(err, &opErr):
			if opErr.Op == "listen" && strings.Contains(opErr.Error(), "address already in use") {
				err = fmt.Errorf("could not open port- check that robotui is not already running (%w)", err)
			}

		case errors.As(err, &pathErr):
			if pathErr.Op == "remove" && strings.Contains(pathErr.Error(), "operation not permitted") {
				err = fmt.Errorf("could not remove file- check that robotui is not already running (%w)", err)
			}
		}
	}

	return err
}
