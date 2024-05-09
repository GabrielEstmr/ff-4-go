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

type UpdateRolloutRequest struct {
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	EnabledAll  bool                   `json:"enabled_all,omitempty"`
	Targets     map[string]interface{} `json:"targets,omitempty"`
}

func (this UpdateRolloutRequest) IsEmpty() bool {
	document := *new(UpdateRolloutRequest)
	return reflect.DeepEqual(this, document)
}

func (this UpdateRolloutRequest) ToDomain(key string) ff_domains.Rollout {
	return *ff_domains.NewRollout(
		key,
		this.Group,
		this.Description,
		this.EnabledAll,
		this.Targets,
	)
}
