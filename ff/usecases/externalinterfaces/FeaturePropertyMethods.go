/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_externalinterfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type FeaturePropertyMethods interface {
	Create(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Update(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	AddValueToProperty(key string, value string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	RemoveValueToProperty(key string, value string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Enable(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Disable(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}
