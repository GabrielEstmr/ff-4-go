package ff_gateways

import (
	ff_resources "baseapplicationgo/main/configs/ff/lib/domains"
	"baseapplicationgo/main/configs/ff/lib/domains/exceptions"
)

type RolloutGateway interface {
	Save(rollout ff_resources.Rollout) (ff_resources.Rollout, ff_domains_exceptions.LibException)
	Update(rollout ff_resources.Rollout) (ff_resources.Rollout, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_resources.Rollout, ff_domains_exceptions.LibException)
}
