package ff_gateways_ws_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type UpdateFeaturePropertyRequest struct {
	Key         string                 `json:"key,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Values      map[string]interface{} `json:"values,omitempty"`
}

func (this UpdateFeaturePropertyRequest) IsEmpty() bool {
	document := *new(UpdateFeaturePropertyRequest)
	return reflect.DeepEqual(this, document)
}

func (this UpdateFeaturePropertyRequest) ToDomain(key string) ff_domains.FeatureProperty {
	return *ff_domains.NewFeatureProperty(
		key,
		this.Group,
		this.Description,
		this.Enabled,
		this.Values,
	)
}