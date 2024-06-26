/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_externalinterfaces

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type RolloutMethods interface {
	Create(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Update(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	AddTargetToRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	RemoveTargetFromRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	EnableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	DisableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	IsEnabledAllOrTargetInRollout(key string, target string) (bool, ff_domains_exceptions.LibException)
}
