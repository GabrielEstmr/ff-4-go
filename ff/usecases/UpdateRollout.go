package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
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
