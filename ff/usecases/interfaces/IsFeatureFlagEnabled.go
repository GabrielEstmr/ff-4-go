package ff_usecases_interfaces

import ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"

type IsFeatureFlagEnabled interface {
	Execute(key string,
	) (bool, ff_domains_exceptions.LibException)
}
