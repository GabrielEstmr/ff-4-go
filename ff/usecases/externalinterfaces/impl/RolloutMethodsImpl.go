/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_externalinterfaces_impl

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
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
