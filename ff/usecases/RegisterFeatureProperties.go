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

type RegisterFeatureProperties struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewRegisterFeatureProperties(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *RegisterFeatureProperties {
	return &RegisterFeatureProperties{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *RegisterFeatureProperties) Execute(
	featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException {
	if featureProperties == nil {
		return nil
	}
	for k, v := range featureProperties {
		persistedFeature, errF := this.featurePropertyGateway.FindById(v.GetKey())
		if errF != nil {
			fmt.Printf("error to find key of name %s\n", k)
			return errF
		}
		if persistedFeature.IsEmpty() {
			_, errS := this.featurePropertyGateway.Save(v)
			if errS != nil {
				fmt.Printf("error to save key of name %s\n", k)
				return errS
			}
			fmt.Printf("feature property with key %s has been saved", v.GetKey())
		}
	}
	return nil
}
