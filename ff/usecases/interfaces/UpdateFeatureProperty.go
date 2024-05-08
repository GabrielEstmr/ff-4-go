package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type UpdateFeatureProperty interface {
	Execute(property ff_domains.FeatureProperty,
	) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}
