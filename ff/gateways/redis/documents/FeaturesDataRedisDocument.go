package ff_gateways_redis_documents

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type FeaturesDataRedisDocument struct {
	Key         string `json:"_id,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Value       bool   `json:"value"`
}

func NewFeaturesDataRedisDocument(featuresData ff_domains.Feature) FeaturesDataRedisDocument {
	return FeaturesDataRedisDocument{
		Key:         featuresData.GetKey(),
		Group:       featuresData.GetGroup(),
		Description: featuresData.GetDescription(),
		Value:       featuresData.GetValue(),
	}
}

func (this FeaturesDataRedisDocument) IsEmpty() bool {
	document := *new(FeaturesDataRedisDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataRedisDocument) ToDomain() ff_domains.Feature {
	if this.IsEmpty() {
		return *new(ff_domains.Feature)
	}
	return *ff_domains.NewFeature(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
	)
}
