package ff_usecases_interfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type RemoveTargetFromRollout interface {
	Execute(key string, target string,
	) (ff_domains.Rollout, ff_domains_exceptions.LibException)
}
