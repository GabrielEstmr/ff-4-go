package ff_gateways

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type FeaturePropertyGateway interface {
	Save(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Update(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}
