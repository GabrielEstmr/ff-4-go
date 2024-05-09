/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type EnableRolloutToAll struct {
	findRolloutById ff_usecases_interfaces.FindRolloutById
	rolloutGateway  ff_gateways.RolloutGateway
}

func NewEnableRolloutToAll(
	findRolloutById ff_usecases_interfaces.FindRolloutById,
	rolloutGateway ff_gateways.RolloutGateway,
) *EnableRolloutToAll {
	return &EnableRolloutToAll{
		findRolloutById: findRolloutById,
		rolloutGateway:  rolloutGateway,
	}
}

func (this *EnableRolloutToAll) Execute(key string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.findRolloutById.Execute(key)
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	persistedRollout.SetEnabledAll(true)
	return this.rolloutGateway.Update(persistedRollout)
}
