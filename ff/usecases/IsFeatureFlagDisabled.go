package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
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
