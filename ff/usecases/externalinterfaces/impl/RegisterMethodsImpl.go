/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_externalinterfaces_impl

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
)

type RegisterMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewRegisterMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *RegisterMethodsImpl {
	return &RegisterMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this *RegisterMethodsImpl) RegisterFeatureFlags(featureFlags ff_domains.FeatureFlags) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterFeatureFlags.Execute(featureFlags)
}

func (this *RegisterMethodsImpl) RegisterFeatureProperties(featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterFeatureProperties.Execute(featureProperties)
}
func (this *RegisterMethodsImpl) RegisterRollouts(rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterRollouts.Execute(rollouts)
}
