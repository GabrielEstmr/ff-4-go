/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_ws_resources

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
	"time"
)

type FeaturePropertyResponse struct {
	Key              string                 `json:"key,omitempty"`
	Group            string                 `json:"group,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Enabled          bool                   `json:"enabled"`
	Values           map[string]interface{} `json:"values,omitempty"`
	CreatedDate      time.Time              `json:"created_date,omitempty"`
	LastModifiedDate time.Time              `json:"last_modified_date,omitempty"`
}

func NewFeaturePropertyResponse(features ff_domains.FeatureProperty) *FeaturePropertyResponse {
	return &FeaturePropertyResponse{
		Key:              features.GetKey(),
		Group:            features.GetGroup(),
		Description:      features.GetDescription(),
		Enabled:          features.GetEnabled(),
		Values:           features.GetValues(),
		CreatedDate:      features.GetCreatedDate(),
		LastModifiedDate: features.GetLastModifiedDate(),
	}
}

func (this FeaturePropertyResponse) IsEmpty() bool {
	document := *new(FeaturePropertyResponse)
	return reflect.DeepEqual(this, document)
}
