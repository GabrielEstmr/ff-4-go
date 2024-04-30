package ff_usecases_interfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RegisterFeatures interface {
	Execute(features ff_domains.Features) ff_domains_exceptions.LibException
}
