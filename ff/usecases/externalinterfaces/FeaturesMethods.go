package ff_usecases_externalinterfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type FeaturesMethods interface {
	Create(feature ff_domains.Feature) (ff_domains.Feature, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	Enable(key string) (ff_domains.Feature, ff_domains_exceptions.LibException)
	Disable(key string) (ff_domains.Feature, ff_domains_exceptions.LibException)
	IsEnabled(key string) (bool, ff_domains_exceptions.LibException)
	IsDisabled(key string) (bool, ff_domains_exceptions.LibException)
}
