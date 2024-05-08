package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type DisableFeatureFlag struct {
	featureGateway   ff_gateways.FeatureFlagsGateway
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey
}

func NewDisableFeatureFlag(
	featureGateway ff_gateways.FeatureFlagsGateway,
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey,
) *DisableFeatureFlag {
	return &DisableFeatureFlag{
		featureGateway:   featureGateway,
		findFeatureByKey: findFeatureByKey,
	}
}

func (this *DisableFeatureFlag) Execute(key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	updatedFeature := feature.CloneAsDisabled()
	return this.featureGateway.Update(updatedFeature)
}
