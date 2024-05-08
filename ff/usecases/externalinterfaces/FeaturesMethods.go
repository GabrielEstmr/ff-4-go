package ff_usecases_externalinterfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FeaturesMethods interface {
	Create(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	Enable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
	Disable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
	IsEnabled(key string) (bool, ff_domains_exceptions.LibException)
	IsDisabled(key string) (bool, ff_domains_exceptions.LibException)
}
