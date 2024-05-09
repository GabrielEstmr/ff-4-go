package ff_usecases

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type EnableFeatureProperty struct {
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById
	featurePropertyGateway  ff_gateways.FeaturePropertyGateway
}

func NewEnableFeatureProperty(
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById,
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *EnableFeatureProperty {
	return &EnableFeatureProperty{
		findFeaturePropertyById: findFeaturePropertyById,
		featurePropertyGateway:  featurePropertyGateway,
	}
}

func (this *EnableFeatureProperty) Execute(key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.findFeaturePropertyById.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	persistedFeature.SetEnabled(true)
	return this.featurePropertyGateway.Update(persistedFeature)
}
