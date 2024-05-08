package ff_gateways_ws_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type FeatureFlagRequest struct {
	Key         string `json:"key,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Value       bool   `json:"value,omitempty"`
}

func (this FeatureFlagRequest) IsEmpty() bool {
	document := *new(FeatureFlagRequest)
	return reflect.DeepEqual(this, document)
}

func (this FeatureFlagRequest) ToDomain() ff_domains.FeatureFlag {
	return *ff_domains.NewFeatureFlag(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
	)
}
