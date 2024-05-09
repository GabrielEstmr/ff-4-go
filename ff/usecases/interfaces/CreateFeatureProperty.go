package ff_usecases_interfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type CreateFeatureProperty interface {
	Execute(property ff_domains.FeatureProperty,
	) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}
