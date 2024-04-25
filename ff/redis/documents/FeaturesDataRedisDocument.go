package ff_redis_documents

import (
	ff_resources "github.com/GabrielEstmr/ff-4-go/ff/resources"
	"reflect"
)

type FeaturesDataRedisDocument struct {
	Key          string `json:"_id,omitempty"`
	Group        string `json:"group,omitempty"`
	Description  string `json:"description,omitempty"`
	DefaultValue bool   `json:"defaultValue"`
}

func NewFeaturesDataRedisDocument(featuresData ff_resources.FeaturesData) FeaturesDataRedisDocument {
	return FeaturesDataRedisDocument{
		Key:          featuresData.GetKey(),
		Group:        featuresData.GetGroup(),
		Description:  featuresData.GetDescription(),
		DefaultValue: featuresData.GetDefaultValue(),
	}
}

func (this FeaturesDataRedisDocument) IsEmpty() bool {
	document := *new(FeaturesDataRedisDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataRedisDocument) ToDomain() ff_resources.FeaturesData {
	if this.IsEmpty() {
		return *new(ff_resources.FeaturesData)
	}
	return *ff_resources.NewFeaturesData(
		this.Key,
		this.Group,
		this.Description,
		this.DefaultValue,
	)
}
