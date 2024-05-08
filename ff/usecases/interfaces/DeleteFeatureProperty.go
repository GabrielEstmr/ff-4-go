package ff_usecases_interfaces

import ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"

type DeleteFeatureProperty interface {
	Execute(key string,
	) ff_domains_exceptions.LibException
}
