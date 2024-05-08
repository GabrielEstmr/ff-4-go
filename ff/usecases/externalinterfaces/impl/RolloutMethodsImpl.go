package ff_usecases_externalinterfaces_impl

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type RolloutMethodsImpl struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewRolloutMethodsImpl(useCaseBeans ff_usecases_beans.UseCaseBeans) *RolloutMethodsImpl {
	return &RolloutMethodsImpl{useCaseBeans: useCaseBeans}
}

func (this RolloutMethodsImpl) Create(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.CreateRollout.Execute(rollout)
}
func (this RolloutMethodsImpl) Update(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.UpdateRollout.Execute(rollout)
}
func (this RolloutMethodsImpl) Delete(key string) ff_domains_exceptions.LibException {
	return this.useCaseBeans.DeleteRollout.Execute(key)
}
func (this RolloutMethodsImpl) FindById(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.FindRolloutById.Execute(key)
}
func (this RolloutMethodsImpl) AddTargetToRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.AddTargetToRollout.Execute(key, target)
}
func (this RolloutMethodsImpl) RemoveTargetFromRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.RemoveTargetFromRollout.Execute(key, target)
}
func (this RolloutMethodsImpl) EnableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.EnableRolloutToAll.Execute(key)
}
func (this RolloutMethodsImpl) DisableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.DisableRolloutToAll.Execute(key)
}
func (this RolloutMethodsImpl) IsEnabledAllOrTargetInRollout(key string, target string) (bool, ff_domains_exceptions.LibException) {
	return this.useCaseBeans.VerifyIsEnabledAllOrTargetInRollout.Execute(key, target)
}
