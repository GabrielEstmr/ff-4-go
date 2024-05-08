package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
)

type DeleteFeatureFlag struct {
	featureGateway ff_gateways.FeatureFlagsGateway
}

func NewDeleteFeatureFlag(featureGateway ff_gateways.FeatureFlagsGateway) *DeleteFeatureFlag {
	return &DeleteFeatureFlag{featureGateway: featureGateway}
}

func (this *DeleteFeatureFlag) Execute(key string) ff_domains_exceptions.LibException {
	return this.featureGateway.Delete(key)
}
