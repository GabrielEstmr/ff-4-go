package ff_gateways

import (
	ff_resources "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FeaturePropertyGateway interface {
	Save(property ff_resources.FeatureProperty) (ff_resources.FeatureProperty, ff_domains_exceptions.LibException)
	Update(property ff_resources.FeatureProperty) (ff_resources.FeatureProperty, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_resources.FeatureProperty, ff_domains_exceptions.LibException)
}
