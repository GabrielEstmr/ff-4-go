/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_redis_documents

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
	"time"
)

type FeaturePropertyDocument struct {
	Key              string                 `json:"_id,omitempty"`
	Group            string                 `json:"group,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Enabled          bool                   `json:"enabled"`
	Values           map[string]interface{} `json:"values,omitempty"`
	CreatedDate      time.Time              `json:"created_date,omitempty"`
	LastModifiedDate time.Time              `json:"last_modified_date,omitempty"`
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
		CreatedDate:      property.GetCreatedDate(),
		LastModifiedDate: property.GetLastModifiedDate(),
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
		this.CreatedDate,
		this.LastModifiedDate,
	)
}

func (this FeaturePropertyDocument) IsEmpty() bool {
	document := *new(FeaturePropertyDocument)
	return reflect.DeepEqual(this, document)
}
