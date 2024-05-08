package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
)

type DeleteFeatureProperty struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewDeleteFeatureProperty(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *DeleteFeatureProperty {
	return &DeleteFeatureProperty{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *DeleteFeatureProperty) Execute(key string,
) ff_domains_exceptions.LibException {
	return this.featurePropertyGateway.Delete(key)
}
