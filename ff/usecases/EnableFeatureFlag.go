package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type EnableFeatureFlag struct {
	featureGateway   ff_gateways.FeatureFlagsGateway
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey
}

func NewEnableFeatureFlag(
	featureGateway ff_gateways.FeatureFlagsGateway,
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey,
) *EnableFeatureFlag {
	return &EnableFeatureFlag{
		featureGateway:   featureGateway,
		findFeatureByKey: findFeatureByKey,
	}
}

func (this *EnableFeatureFlag) Execute(key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	updatedFeature := feature.CloneAsEnabled()
	return this.featureGateway.Update(updatedFeature)
}
