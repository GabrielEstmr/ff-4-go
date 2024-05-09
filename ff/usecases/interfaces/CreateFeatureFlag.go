package ff_usecases_interfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type CreateFeatureFlag interface {
	Execute(featureFlag ff_domains.FeatureFlag,
	) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
}
