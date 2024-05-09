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

type FeatureFlagResponse struct {
	Key              string    `json:"key,omitempty"`
	Group            string    `json:"group,omitempty"`
	Description      string    `json:"description,omitempty"`
	Value            bool      `json:"value"`
	CreatedDate      time.Time `json:"created_date,omitempty"`
	LastModifiedDate time.Time `json:"last_modified_date,omitempty"`
}

func NewFeatureResponse(
	featureFlag ff_domains.FeatureFlag,
) *FeatureFlagResponse {
	return &FeatureFlagResponse{
		Key:              featureFlag.GetKey(),
		Group:            featureFlag.GetGroup(),
		Description:      featureFlag.GetDescription(),
		Value:            featureFlag.GetValue(),
		CreatedDate:      featureFlag.GetCreatedDate(),
		LastModifiedDate: featureFlag.GetLastModifiedDate(),
	}
}

func (this FeatureFlagResponse) IsEmpty() bool {
	document := *new(FeatureFlagResponse)
	return reflect.DeepEqual(this, document)
}
