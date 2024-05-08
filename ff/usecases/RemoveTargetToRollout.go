package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type RemoveTargetFromRollout struct {
	findRolloutById ff_usecases_interfaces.FindRolloutById
	rolloutGateway  ff_gateways.RolloutGateway
}

func NewRemoveTargetFromRollout(
	findRolloutById ff_usecases_interfaces.FindRolloutById,
	rolloutGateway ff_gateways.RolloutGateway,
) *RemoveTargetFromRollout {
	return &RemoveTargetFromRollout{
		findRolloutById: findRolloutById,
		rolloutGateway:  rolloutGateway,
	}
}

func (this *RemoveTargetFromRollout) Execute(key string, target string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.findRolloutById.Execute(key)
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	persistedRollout.RemoveTarget(target)
	return this.rolloutGateway.Update(persistedRollout)
}
