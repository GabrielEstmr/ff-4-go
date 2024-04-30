package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type EnableFeature struct {
	featureGateway   ff_gateways.FeaturesGateway
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey
}

func NewEnableFeature(
	featureGateway ff_gateways.FeaturesGateway,
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey,
) *EnableFeature {
	return &EnableFeature{
		featureGateway:   featureGateway,
		findFeatureByKey: findFeatureByKey,
	}
}

func (this *EnableFeature) Execute(key string,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return *new(ff_domains.Feature), errF
	}

	updatedFeature := feature.CloneAsEnabled()
	return this.featureGateway.Update(updatedFeature)
}
