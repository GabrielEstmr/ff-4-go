package ff_gateways_redis_documents

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
	"time"
)

type FeatureFlagDocument struct {
	Key              string    `json:"_id,omitempty"`
	Group            string    `json:"group,omitempty"`
	Description      string    `json:"description,omitempty"`
	Value            bool      `json:"value"`
	CreatedDate      time.Time `json:"created_date,omitempty"`
	LastModifiedDate time.Time `json:"last_modified_date,omitempty"`
}

func NewFeatureFlagDocument(featureFlag ff_domains.FeatureFlag) FeatureFlagDocument {
	return FeatureFlagDocument{
		Key:              featureFlag.GetKey(),
		Group:            featureFlag.GetGroup(),
		Description:      featureFlag.GetDescription(),
		Value:            featureFlag.GetValue(),
		CreatedDate:      featureFlag.GetCreatedDate(),
		LastModifiedDate: featureFlag.GetLastModifiedDate(),
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
		this.CreatedDate,
		this.LastModifiedDate,
	)
}
