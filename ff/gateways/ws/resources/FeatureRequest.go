package ff_gateways_ws_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type FeatureRequest struct {
	Key         string `json:"key,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Value       bool   `json:"value,omitempty"`
}

func (this FeatureRequest) IsEmpty() bool {
	document := *new(FeatureRequest)
	return reflect.DeepEqual(this, document)
}

func (this FeatureRequest) ToDomain() ff_domains.Feature {
	return *ff_domains.NewFeature(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
	)
}
