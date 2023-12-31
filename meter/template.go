package meter

import (
	"github.com/thommyho/robotui/api"
	"github.com/thommyho/robotui/util/templates"
)

func init() {
	registry.Add("template", NewMeterFromTemplateConfig)
}

func NewMeterFromTemplateConfig(other map[string]interface{}) (api.Meter, error) {
	instance, err := templates.RenderInstance(templates.Meter, other)

	var res api.Meter
	if err == nil {
		res, err = NewFromConfig(instance.Type, instance.Other)
	}

	return res, err
}
