package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type CreateRollout interface {
	Execute(rollout ff_domains.Rollout,
	) (ff_domains.Rollout, ff_domains_exceptions.LibException)
}
