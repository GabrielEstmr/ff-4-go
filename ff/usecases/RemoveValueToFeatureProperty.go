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

type RemoveValueToFeatureProperty struct {
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById
	featurePropertyGateway  ff_gateways.FeaturePropertyGateway
}

func NewRemoveValueToFeatureProperty(
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById,
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *RemoveValueToFeatureProperty {
	return &RemoveValueToFeatureProperty{
		findFeaturePropertyById: findFeaturePropertyById,
		featurePropertyGateway:  featurePropertyGateway,
	}
}

func (this *RemoveValueToFeatureProperty) Execute(key string, value string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.findFeaturePropertyById.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	persistedFeature.RemoveValue(value)
	return this.featurePropertyGateway.Update(persistedFeature)
}
