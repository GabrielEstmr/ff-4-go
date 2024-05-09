package ff_usecases_interfaces

import ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"

type DeleteFeatureFlag interface {
	Execute(key string) ff_domains_exceptions.LibException
}
