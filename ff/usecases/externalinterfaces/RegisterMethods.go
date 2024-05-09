package ff_usecases_externalinterfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type RegisterMethodsImpl interface {
	RegisterFeatureFlags(featuresFlags ff_domains.FeatureFlags) ff_domains_exceptions.LibException
	RegisterFeatureProperties(featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException
	RegisterRollouts(rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException
}
