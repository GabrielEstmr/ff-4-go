package ff_usecases_externalinterfaces_impl

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type RegisterMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewRegisterMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *RegisterMethodsImpl {
	return &RegisterMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this *RegisterMethodsImpl) RegisterFeatureFlags(featureFlags ff_domains.FeatureFlags) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterFeatureFlags.Execute(featureFlags)
}

func (this *RegisterMethodsImpl) RegisterFeatureProperties(featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterFeatureProperties.Execute(featureProperties)
}
func (this *RegisterMethodsImpl) RegisterRollouts(rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException {
	return this.useCaseBeans.RegisterRollouts.Execute(rollouts)
}
