package ff_gateways

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FeaturesGateway interface {
	Save(feature ff_domains.Feature) (ff_domains.Feature, ff_domains_exceptions.LibException)
	Update(feature ff_domains.Feature) (ff_domains.Feature, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.Feature, ff_domains_exceptions.LibException)
}
