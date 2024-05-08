package ff_usecases_interfaces

import ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"

type VerifyIsEnabledAllOrTargetInRollout interface {
	Execute(key string, target string,
	) (bool, ff_domains_exceptions.LibException)
}
