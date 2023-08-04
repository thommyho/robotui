package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/core"
	"github.com/thommyho/robotui/util"
)

const sample = `
loadpoints:
- mode: off
`

func TestYamlOff(t *testing.T) {
	var conf config
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(sample)); err != nil {
		t.Error(err)
	}

	if err := viper.UnmarshalExact(&conf); err != nil {
		t.Error(err)
	}

	var lp core.Loadpoint
	if err := util.DecodeOther(conf.Loadpoints[0], &lp); err != nil {
		t.Error(err)
	}

	if lp.Mode != api.ModeOff {
		t.Errorf("expected `off`, got %s", lp.Mode)
	}
}
