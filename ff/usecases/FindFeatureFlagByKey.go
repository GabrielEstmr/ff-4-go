/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type FindFeatureFlagByKey struct {
	featureGateway ff_gateways.FeatureFlagsGateway
}

func NewFindFeatureFlagByKey(featureGateway ff_gateways.FeatureFlagsGateway) *FindFeatureFlagByKey {
	return &FindFeatureFlagByKey{featureGateway: featureGateway}
}

func (this *FindFeatureFlagByKey) Execute(key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featureGateway.FindById(key)
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	if persistedFeature.IsEmpty() {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("feature with key %s not found.", key))
	}

	return persistedFeature, nil
}
