package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
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
