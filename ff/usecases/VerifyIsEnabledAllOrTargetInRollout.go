package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
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
