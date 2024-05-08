package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type VerifyIsEnabledAllOrTargetInRollout struct {
	findRolloutById ff_usecases_interfaces.FindRolloutById
}

func NewVerifyIsEnabledAllOrTargetInRollout(
	findRolloutById ff_usecases_interfaces.FindRolloutById,
) *VerifyIsEnabledAllOrTargetInRollout {
	return &VerifyIsEnabledAllOrTargetInRollout{
		findRolloutById: findRolloutById,
	}
}

func (this *VerifyIsEnabledAllOrTargetInRollout) Execute(key string, target string,
) (bool, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.findRolloutById.Execute(key)
	if errF != nil {
		return false, errF
	}
	return persistedRollout.IsEnabledAllOrTargetInRollout(target), nil
}
