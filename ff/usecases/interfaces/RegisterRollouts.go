package ff_usecases_interfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type RegisterRollouts interface {
	Execute(
		rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException
}
