package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RegisterRollouts interface {
	Execute(
		rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException
}
