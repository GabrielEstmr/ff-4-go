package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
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
