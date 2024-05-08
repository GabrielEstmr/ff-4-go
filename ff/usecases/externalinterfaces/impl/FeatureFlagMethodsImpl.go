package ff_usecases_externalinterfaces_impl

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type FeatureFlagMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeatureFlagMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *FeatureFlagMethodsImpl {
	return &FeatureFlagMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this *FeatureFlagMethodsImpl) Create(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.CreateFeatureFlag.Execute(featureFlag)
}
func (this *FeatureFlagMethodsImpl) Delete(key string) ff_domains_exceptions.LibException {
	return this.useCaseBeans.DeleteFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) Enable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.EnableFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) Disable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.DisableFeatureFlag.Execute(key)
}
func (this *FeatureFlagMethodsImpl) IsEnabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureFlagEnabled.Execute(key)
}
func (this *FeatureFlagMethodsImpl) IsDisabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureFlagEnabled.Execute(key)
}