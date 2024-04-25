package ff_mongo_documents

import (
	ff_resources "github.com/GabrielEstmr/ff-4-go/ff/resources"
	"reflect"
)

type FeaturesDataDocument struct {
	Key          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Group        string `json:"group,omitempty" bson:"group,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
	DefaultValue bool   `json:"defaultValue" bson:"defaultValue"`
}

func NewFeaturesDataDocument(featuresData ff_resources.FeaturesData) FeaturesDataDocument {
	return FeaturesDataDocument{
		Key:          featuresData.GetKey(),
		Group:        featuresData.GetGroup(),
		Description:  featuresData.GetDescription(),
		DefaultValue: featuresData.GetDefaultValue(),
	}
}

func (this FeaturesDataDocument) IsEmpty() bool {
	document := *new(FeaturesDataDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataDocument) ToDomain() ff_resources.FeaturesData {
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

func (this FeaturesDataDocument) IsEnabled() bool {
	return this.DefaultValue == true
}

func (this FeaturesDataDocument) IsDisabled() bool {
	return this.DefaultValue == false
}
