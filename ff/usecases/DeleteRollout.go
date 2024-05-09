/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type DeleteRollout struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewDeleteRollout(rolloutGateway ff_gateways.RolloutGateway) *DeleteRollout {
	return &DeleteRollout{rolloutGateway: rolloutGateway}
}

func (this *DeleteRollout) Execute(key string,
) ff_domains_exceptions.LibException {
	return this.rolloutGateway.Delete(key)
}
