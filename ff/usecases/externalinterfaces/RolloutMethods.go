package ff_usecases_externalinterfaces

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RolloutMethods interface {
	Create(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Update(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	//FindByFilter(filter ff_domains.RolloutFilter) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	AddTargetToRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	RemoveTargetFromRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	EnableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	DisableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	IsEnabledAllOrTargetInRollout(key string, target string) (bool, ff_domains_exceptions.LibException)
}
