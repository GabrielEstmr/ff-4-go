package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type CreateRollout struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewCreateRollout(rolloutGateway ff_gateways.RolloutGateway) *CreateRollout {
	return &CreateRollout{rolloutGateway: rolloutGateway}
}

func (this *CreateRollout) Execute(rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	persistedRollout, errF := this.rolloutGateway.FindById(rollout.GetKey())
	if errF != nil {
		return *new(ff_domains.Rollout), errF
	}

	if !persistedRollout.IsEmpty() {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewConflictExceptionSglMsg(
				fmt.Sprintf("Feature with key %s already exists.", rollout.GetKey()))
	}

	return this.rolloutGateway.Save(rollout)
}
