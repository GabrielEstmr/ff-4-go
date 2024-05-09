package ff_usecases_interfaces

import ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"

type IsFeatureFlagEnabled interface {
	Execute(key string,
	) (bool, ff_domains_exceptions.LibException)
}
