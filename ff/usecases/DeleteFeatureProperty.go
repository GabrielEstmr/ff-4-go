/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type DeleteFeatureProperty struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewDeleteFeatureProperty(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *DeleteFeatureProperty {
	return &DeleteFeatureProperty{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *DeleteFeatureProperty) Execute(key string,
) ff_domains_exceptions.LibException {
	return this.featurePropertyGateway.Delete(key)
}
