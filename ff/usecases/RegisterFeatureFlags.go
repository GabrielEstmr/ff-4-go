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

type RegisterFeatureFlags struct {
	featureGateway ff_gateways.FeatureFlagsGateway
}

func NewRegisterFeatureFlags(featureGateway ff_gateways.FeatureFlagsGateway) *RegisterFeatureFlags {
	return &RegisterFeatureFlags{featureGateway: featureGateway}
}

func (this *RegisterFeatureFlags) Execute(featureFlags ff_domains.FeatureFlags) ff_domains_exceptions.LibException {
	if featureFlags == nil {
		return nil
	}
	for k, v := range featureFlags {
		persistedFeature, errF := this.featureGateway.FindById(v.GetKey())
		if errF != nil {
			fmt.Printf("error to find key of name %s\n", k)
			return errF
		}
		if persistedFeature.IsEmpty() {
			_, errS := this.featureGateway.Save(v)
			if errS != nil {
				fmt.Printf("error to save key of name %s\n", k)
				return errS
			}
			fmt.Printf("feature with key %s has been saved", v.GetKey())
		}
	}
	return nil
}
