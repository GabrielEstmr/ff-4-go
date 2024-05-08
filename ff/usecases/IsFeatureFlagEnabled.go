package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type IsFeatureFlagEnabled struct {
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey
}

func NewIsFeatureFlagEnabled(findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey) *IsFeatureFlagEnabled {
	return &IsFeatureFlagEnabled{findFeatureByKey: findFeatureByKey}
}

func (this *IsFeatureFlagEnabled) Execute(key string,
) (bool, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return false, errF
	}

	return feature.IsEnabled(), nil
}
