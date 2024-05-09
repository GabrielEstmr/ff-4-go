/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type VerifyIsEnabledAllOrTargetInRollout struct {
	findRolloutById ff_usecases_interfaces.FindRolloutById
}

func NewVerifyIsEnabledAllOrTargetInRollout(
	findRolloutById ff_usecases_interfaces.FindRolloutById,
) *VerifyIsEnabledAllOrTargetInRollout {
	return &VerifyIsEnabledAllOrTargetInRollout{
		findRolloutById: findRolloutById,
	}
}

func (this *VerifyIsEnabledAllOrTargetInRollout) Execute(key string, target string,
) (bool, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.findRolloutById.Execute(key)
	if errF != nil {
		return false, errF
	}
	return persistedRollout.IsEnabledAllOrTargetInRollout(target), nil
}
