package ff_usecases_externalinterfaces_impl

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type FeaturesMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeaturesMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *FeaturesMethodsImpl {
	return &FeaturesMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this *FeaturesMethodsImpl) Create(feature ff_domains.Feature) (ff_domains.Feature, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.CreateFeature.Execute(feature)
}
func (this *FeaturesMethodsImpl) Delete(key string) ff_domains_exceptions.LibException {
	return this.useCaseBeans.DeleteFeature.Execute(key)
}
func (this *FeaturesMethodsImpl) Enable(key string) (ff_domains.Feature, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.EnableFeature.Execute(key)
}
func (this *FeaturesMethodsImpl) Disable(key string) (ff_domains.Feature, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.DisableFeature.Execute(key)
}
func (this *FeaturesMethodsImpl) IsEnabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureEnabled.Execute(key)
}
func (this *FeaturesMethodsImpl) IsDisabled(key string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.IsFeatureEnabled.Execute(key)
}
