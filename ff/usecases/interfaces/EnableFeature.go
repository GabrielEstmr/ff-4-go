package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type EnableFeature interface {
	Execute(key string,
	) (ff_domains.Feature, ff_domains_exceptions.LibException)
}
