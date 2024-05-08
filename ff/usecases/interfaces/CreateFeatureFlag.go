package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type CreateFeatureFlag interface {
	Execute(featureFlag ff_domains.FeatureFlag,
	) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
}
