package ff_usecases

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type AddValueToFeatureProperty struct {
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById
	featurePropertyGateway  ff_gateways.FeaturePropertyGateway
}

func NewAddValueToFeatureProperty(
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById,
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *AddValueToFeatureProperty {
	return &AddValueToFeatureProperty{
		findFeaturePropertyById: findFeaturePropertyById,
		featurePropertyGateway:  featurePropertyGateway,
	}
}

func (this *AddValueToFeatureProperty) Execute(key string, value string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.findFeaturePropertyById.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	persistedFeature.AddValue(value)
	return this.featurePropertyGateway.Update(persistedFeature)
}
