package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type DisableFeature struct {
	featureGateway   ff_gateways.FeaturesGateway
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey
}

func NewDisableFeature(
	featureGateway ff_gateways.FeaturesGateway,
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey,
) *DisableFeature {
	return &DisableFeature{
		featureGateway:   featureGateway,
		findFeatureByKey: findFeatureByKey,
	}
}

func (this *DisableFeature) Execute(key string,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return *new(ff_domains.Feature), errF
	}

	updatedFeature := feature.CloneAsDisabled()
	return this.featureGateway.Update(updatedFeature)
}
