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

type FeatureFlagMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeatureFlagMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *FeatureFlagMethodsImpl {
	return &FeatureFlagMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this *FeatureFlagMethodsImpl) Create(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.CreateFeatureFlag.Execute(featureFlag)
}
func (this *FeatureFlagMethodsImpl) Delete(key string) ff_domains_exceptions.LibException {
	return this.useCaseBeans.DeleteFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) Enable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.EnableFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) Disable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.DisableFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) IsEnabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureFlagEnabled.Execute(key)
}
func (this *FeatureFlagMethodsImpl) IsDisabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureFlagEnabled.Execute(key)
}
