package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type AddTargetToRollout struct {
	findRolloutById ff_usecases_interfaces.FindRolloutById
	rolloutGateway  ff_gateways.RolloutGateway
}

func NewAddTargetToRollout(
	findRolloutById ff_usecases_interfaces.FindRolloutById,
	rolloutGateway ff_gateways.RolloutGateway,
) *AddTargetToRollout {
	return &AddTargetToRollout{
		findRolloutById: findRolloutById,
		rolloutGateway:  rolloutGateway,
	}
}

func (this *AddTargetToRollout) Execute(key string, target string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.findRolloutById.Execute(key)
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	persistedRollout.AddTarget(target)
	return this.rolloutGateway.Update(persistedRollout)
}
