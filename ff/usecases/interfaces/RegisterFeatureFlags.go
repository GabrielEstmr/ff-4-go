package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RegisterFeatureFlags interface {
	Execute(featureFlags ff_domains.FeatureFlags) ff_domains_exceptions.LibException
}
