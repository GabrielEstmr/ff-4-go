package ff_gateways_mongo_documents

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type FeatureFlagDocument struct {
	Key              string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Group            string             `json:"group,omitempty" bson:"group,omitempty"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	Value            bool               `json:"value" bson:"value"`
	CreatedDate      primitive.DateTime `json:"created_date,omitempty" bson:"created_date"`
	LastModifiedDate primitive.DateTime `json:"last_modified_date,omitempty" bson:"last_modified_date"`
}

func NewFeatureFlagDocument(featureFlag ff_domains.FeatureFlag) FeatureFlagDocument {
	return FeatureFlagDocument{
		Key:              featureFlag.GetKey(),
		Group:            featureFlag.GetGroup(),
		Description:      featureFlag.GetDescription(),
		Value:            featureFlag.GetValue(),
		CreatedDate:      primitive.NewDateTimeFromTime(featureFlag.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(featureFlag.GetLastModifiedDate()),
	}
}

func (this FeatureFlagDocument) IsEmpty() bool {
	document := *new(FeatureFlagDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeatureFlagDocument) ToDomain() ff_domains.FeatureFlag {
	if this.IsEmpty() {
		return *new(ff_domains.FeatureFlag)
	}
	return *ff_domains.NewFeatureFlagAllArgs(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}

func (this FeatureFlagDocument) IsEnabled() bool {
	return this.Value == true
}

func (this FeatureFlagDocument) IsDisabled() bool {
	return this.Value == false
}
