package ff_gateways_ws_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type FeatureResponse struct {
	Key         string `json:"key,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Value       bool   `json:"value"`
}

func NewFeatureResponse(
	feature ff_domains.Feature,
) *FeatureResponse {
	return &FeatureResponse{
		Key:         feature.GetKey(),
		Group:       feature.GetGroup(),
		Description: feature.GetDescription(),
		Value:       feature.GetValue(),
	}
}

func (this FeatureResponse) IsEmpty() bool {
	document := *new(FeatureResponse)
	return reflect.DeepEqual(this, document)
}
