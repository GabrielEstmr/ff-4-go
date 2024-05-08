package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RegisterFeatureProperties interface {
	Execute(
		featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException
}
