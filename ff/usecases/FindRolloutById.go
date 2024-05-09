/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type FindRolloutById struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewFindRolloutById(rolloutGateway ff_gateways.RolloutGateway) *FindRolloutById {
	return &FindRolloutById{rolloutGateway: rolloutGateway}
}

func (this *FindRolloutById) Execute(key string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.rolloutGateway.FindById(key)
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	if persistedRollout.IsEmpty() {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("rollout with key %s not found.", key))
	}

	return persistedRollout, nil
}
