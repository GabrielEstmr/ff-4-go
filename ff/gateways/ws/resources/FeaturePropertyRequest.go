package ff_gateways_ws_resources

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
)

type FeaturePropertyRequest struct {
	Key         string                 `json:"key,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Values      map[string]interface{} `json:"values,omitempty"`
}

func (this FeaturePropertyRequest) IsEmpty() bool {
	document := *new(FeaturePropertyRequest)
	return reflect.DeepEqual(this, document)
}

func (this FeaturePropertyRequest) ToDomain() ff_domains.FeatureProperty {
	return *ff_domains.NewFeatureProperty(
		this.Key,
		this.Group,
		this.Description,
		this.Enabled,
		this.Values,
	)
}
