package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type IsFeatureEnabled struct {
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey
}

func NewIsFeatureEnabled(findFeatureByKey ff_usecases_interfaces.FindFeatureByKey) *IsFeatureEnabled {
	return &IsFeatureEnabled{findFeatureByKey: findFeatureByKey}
}

func (this *IsFeatureEnabled) Execute(key string,
) (bool, ff_domains_exceptions.LibException) {

	feature, errF := this.findFeatureByKey.Execute(key)
	if errF != nil {
		return false, errF
	}

	return feature.IsEnabled(), nil
}
