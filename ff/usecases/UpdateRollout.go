package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type UpdateRollout struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewUpdateRollout(rolloutGateway ff_gateways.RolloutGateway) *UpdateRollout {
	return &UpdateRollout{rolloutGateway: rolloutGateway}
}

func (this *UpdateRollout) Execute(rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.rolloutGateway.FindById(rollout.GetKey())
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	if persistedRollout.IsEmpty() {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("rollout with key %s do not exists.", rollout.GetKey()))
	}

	return this.rolloutGateway.Update(rollout)
}
