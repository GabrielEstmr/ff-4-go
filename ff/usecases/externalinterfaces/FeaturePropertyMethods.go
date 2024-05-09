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
