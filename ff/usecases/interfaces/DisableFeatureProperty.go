package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type DisableFeatureProperty interface {
	Execute(key string,
	) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}
