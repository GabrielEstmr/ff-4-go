package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type IsFeatureFlagDisabled struct {
	isFeatureEnabled ff_usecases_interfaces.IsFeatureFlagEnabled
}

func NewIsFeatureFlagDisabled(isFeatureEnabled ff_usecases_interfaces.IsFeatureFlagEnabled) *IsFeatureFlagDisabled {
	return &IsFeatureFlagDisabled{isFeatureEnabled: isFeatureEnabled}
}

func (this *IsFeatureFlagDisabled) Execute(key string,
) (bool, ff_domains_exceptions.LibException) {
	isEnabled, err := this.isFeatureEnabled.Execute(key)
	return !isEnabled, err
}
