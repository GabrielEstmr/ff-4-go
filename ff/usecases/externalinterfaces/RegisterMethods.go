package ff_usecases_externalinterfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RegisterMethods interface {
	RegisterFeature(features ff_domains.Features) ff_domains_exceptions.LibException
}
