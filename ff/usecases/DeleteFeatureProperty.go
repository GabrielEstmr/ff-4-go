package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
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
