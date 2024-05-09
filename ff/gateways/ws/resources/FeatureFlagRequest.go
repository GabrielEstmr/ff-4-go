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
)

type FeatureFlagRequest struct {
	Key         string `json:"key,omitempty"`
	Group       string `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Value       bool   `json:"value,omitempty"`
}

func (this FeatureFlagRequest) IsEmpty() bool {
	document := *new(FeatureFlagRequest)
	return reflect.DeepEqual(this, document)
}

func (this FeatureFlagRequest) ToDomain() ff_domains.FeatureFlag {
	return *ff_domains.NewFeatureFlag(
		this.Key,
		this.Group,
		this.Description,
		this.Value,
	)
}
