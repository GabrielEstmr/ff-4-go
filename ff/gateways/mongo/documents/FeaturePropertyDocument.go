package ff_gateways_mongo_documents

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type FeaturePropertyDocument struct {
	Key              string                 `json:"_id,omitempty" bson:"_id,omitempty"`
	Group            string                 `json:"group,omitempty" bson:"group,omitempty"`
	Description      string                 `json:"description,omitempty" bson:"description,omitempty"`
	Enabled          bool                   `json:"enabled" bson:"enabled"`
	Values           map[string]interface{} `json:"values,omitempty" bson:"values,omitempty"`
	CreatedDate      primitive.DateTime     `json:"created_date,omitempty" bson:"created_date"`
	LastModifiedDate primitive.DateTime     `json:"last_modified_date,omitempty" bson:"last_modified_date"`
}

func NewFeaturePropertyDocument(
	property ff_domains.FeatureProperty,
) *FeaturePropertyDocument {
	return &FeaturePropertyDocument{
		Key:              property.GetKey(),
		Group:            property.GetGroup(),
		Description:      property.GetDescription(),
		Enabled:          property.GetEnabled(),
		Values:           property.GetValues(),
		CreatedDate:      primitive.NewDateTimeFromTime(property.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(property.GetLastModifiedDate()),
	}
}

func (this FeaturePropertyDocument) ToDomain() ff_domains.FeatureProperty {
	if this.IsEmpty() {
		return *new(ff_domains.FeatureProperty)
	}
	return *ff_domains.NewFeaturePropertyAllArgs(
		this.Key,
		this.Group,
		this.Description,
		this.Enabled,
		this.Values,
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}

func (this FeaturePropertyDocument) IsEmpty() bool {
	document := *new(FeaturePropertyDocument)
	return reflect.DeepEqual(this, document)
}
