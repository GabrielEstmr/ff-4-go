package ff_gateways

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
)

type RolloutGateway interface {
	Save(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Update(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
	Delete(key string) ff_domains_exceptions.LibException
	FindById(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
}
