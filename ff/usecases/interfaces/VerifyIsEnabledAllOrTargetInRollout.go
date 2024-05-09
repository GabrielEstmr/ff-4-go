package ff_usecases_interfaces

import ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"

type VerifyIsEnabledAllOrTargetInRollout interface {
	Execute(key string, target string,
	) (bool, ff_domains_exceptions.LibException)
}
