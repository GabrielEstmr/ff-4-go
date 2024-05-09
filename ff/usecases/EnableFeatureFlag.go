/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type EnableFeatureFlag struct {
	featureGateway   ff_gateways.FeatureFlagsGateway
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey
}

func NewEnableFeatureFlag(
	featureGateway ff_gateways.FeatureFlagsGateway,
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey,
) *EnableFeatureFlag {
	return &EnableFeatureFlag{
		featureGateway:   featureGateway,
		findFeatureByKey: findFeatureByKey,
	}
}

func (this *EnableFeatureFlag) Execute(key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	updatedFeature := feature.CloneAsEnabled()
	return this.featureGateway.Update(updatedFeature)
}
