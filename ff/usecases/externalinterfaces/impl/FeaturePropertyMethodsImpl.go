package ff_usecases_externalinterfaces_impl

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
)

type FeaturePropertyMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeaturePropertyMethodsImpl(
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *FeaturePropertyMethodsImpl {
	return &FeaturePropertyMethodsImpl{
		useCaseBeans: useCaseBeans,
	}
}

func (this FeaturePropertyMethodsImpl) Create(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.CreateFeatureProperty.Execute(property)
}

func (this FeaturePropertyMethodsImpl) Update(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.UpdateFeatureProperty.Execute(property)
}

func (this FeaturePropertyMethodsImpl) Delete(key string) ff_domains_exceptions.LibException {
	return this.useCaseBeans.DeleteFeatureProperty.Execute(key)
}

func (this FeaturePropertyMethodsImpl) FindById(key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.FindFeaturePropertyById.Execute(key)
}

func (this FeaturePropertyMethodsImpl) AddValueToProperty(key string, value string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.AddValueToFeatureProperty.Execute(key, value)
}

func (this FeaturePropertyMethodsImpl) RemoveValueToProperty(key string, value string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.RemoveValueToFeatureProperty.Execute(key, value)
}

func (this FeaturePropertyMethodsImpl) Enable(key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.EnableFeatureProperty.Execute(key)
}

func (this FeaturePropertyMethodsImpl) Disable(key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.DisableFeatureProperty.Execute(key)
}
