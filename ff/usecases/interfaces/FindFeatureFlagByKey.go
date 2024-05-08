package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FindFeatureFlagByKey interface {
	Execute(key string,
	) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
}
