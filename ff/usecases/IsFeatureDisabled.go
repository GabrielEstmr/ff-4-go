package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type IsFeatureDisabled struct {
	isFeatureEnabled ff_usecases_interfaces.IsFeatureEnabled
}

func NewIsFeatureDisabled(isFeatureEnabled ff_usecases_interfaces.IsFeatureEnabled) *IsFeatureDisabled {
	return &IsFeatureDisabled{isFeatureEnabled: isFeatureEnabled}
}

func (this *IsFeatureDisabled) Execute(key string,
) (bool, ff_domains_exceptions.LibException) {
	isEnabled, err := this.isFeatureEnabled.Execute(key)
	return !isEnabled, err
}
