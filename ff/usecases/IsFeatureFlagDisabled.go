/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type IsFeatureFlagDisabled struct {
	isFeatureEnabled ff_usecases_interfaces.IsFeatureFlagEnabled
}

func NewIsFeatureFlagDisabled(isFeatureEnabled ff_usecases_interfaces.IsFeatureFlagEnabled) *IsFeatureFlagDisabled {
	return &IsFeatureFlagDisabled{isFeatureEnabled: isFeatureEnabled}
}

func (this *IsFeatureFlagDisabled) Execute(key string,
) (bool, ff_domains_exceptions.LibException) {
	isEnabled, err := this.isFeatureEnabled.Execute(key)
	return !isEnabled, err
}
