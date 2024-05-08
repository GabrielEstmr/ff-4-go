package ff_gateways

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FeatureFlagsGateway interface {
	Save(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
	Update(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
}
