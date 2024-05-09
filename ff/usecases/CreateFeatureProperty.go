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

type CreateFeatureProperty struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewCreateFeatureProperty(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *CreateFeatureProperty {
	return &CreateFeatureProperty{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *CreateFeatureProperty) Execute(property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedProperty, errF := this.featurePropertyGateway.FindById(property.GetKey())
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	if !persistedProperty.IsEmpty() {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewConflictExceptionSglMsg(
				fmt.Sprintf("Feature property with key %s already exists.", property.GetKey()))
	}

	return this.featurePropertyGateway.Save(property)
}
