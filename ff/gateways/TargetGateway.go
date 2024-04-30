package ff_gateways

import (
	ff_resources "baseapplicationgo/main/configs/ff/lib/domains"
	"baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RolloutMethods interface {
	Save(Rollout ff_resources.Rollout) (ff_resources.Rollout, ff_domains_exceptions.LibException)
	Update(Rollout ff_resources.Rollout) (ff_resources.Rollout, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_resources.Rollout, ff_domains_exceptions.LibException)
	FindByFilter(filter ff_resources.RolloutFilter) (ff_resources.Rollout, ff_domains_exceptions.LibException)
}
