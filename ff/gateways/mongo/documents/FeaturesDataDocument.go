package ff_gateways_mongo_documents

import (
	ff_resources "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
)

type FeaturesDataDocument struct {
	Key         string `json:"_id,omitempty" bson:"_id,omitempty"`
	Group       string `json:"group,omitempty" bson:"group,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Value       bool   `json:"value" bson:"value"`
}

func NewFeaturesDataDocument(featuresData ff_resources.Feature) FeaturesDataDocument {
	return FeaturesDataDocument{
		Key:         featuresData.GetKey(),
		Group:       featuresData.GetGroup(),
		Description: featuresData.GetDescription(),
		Value:       featuresData.GetValue(),
	}
}

func (this FeaturesDataDocument) IsEmpty() bool {
	document := *new(FeaturesDataDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataDocument) ToDomain() ff_resources.Feature {
	if this.IsEmpty() {
		return *new(ff_resources.Feature)
	}
	return *ff_resources.NewFeature(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
	)
}

func (this FeaturesDataDocument) IsEnabled() bool {
	return this.Value == true
}

func (this FeaturesDataDocument) IsDisabled() bool {
	return this.Value == false
}
