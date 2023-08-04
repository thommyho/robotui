package vehicle

import (
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/util/templates"
)

func init() {
	registry.Add("template", NewVehicleFromTemplateConfig)
}

func NewVehicleFromTemplateConfig(other map[string]interface{}) (api.Vehicle, error) {
	instance, err := templates.RenderInstance(templates.Vehicle, other)

	var res api.Vehicle
	if err == nil {
		res, err = NewFromConfig(instance.Type, instance.Other)
	}

	return res, err
}
