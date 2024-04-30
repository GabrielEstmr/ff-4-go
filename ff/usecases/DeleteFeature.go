package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
)

type DeleteFeature struct {
	featureGateway ff_gateways.FeaturesGateway
}

func NewDeleteFeature(featureGateway ff_gateways.FeaturesGateway) *DeleteFeature {
	return &DeleteFeature{featureGateway: featureGateway}
}

func (this *DeleteFeature) Execute(key string) ff_domains_exceptions.LibException {
	return this.featureGateway.Delete(key)
}
