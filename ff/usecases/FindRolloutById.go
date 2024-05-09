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
