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

type UpdateFeaturePropertyRequest struct {
	Key         string                 `json:"key,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Values      map[string]interface{} `json:"values,omitempty"`
}

func (this UpdateFeaturePropertyRequest) IsEmpty() bool {
	document := *new(UpdateFeaturePropertyRequest)
	return reflect.DeepEqual(this, document)
}

func (this UpdateFeaturePropertyRequest) ToDomain(key string) ff_domains.FeatureProperty {
	return *ff_domains.NewFeatureProperty(
		key,
		this.Group,
		this.Description,
		this.Enabled,
		this.Values,
	)
}
