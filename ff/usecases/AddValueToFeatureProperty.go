package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
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
