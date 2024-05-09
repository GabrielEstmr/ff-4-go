package ff_usecases

import (
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type DeleteRollout struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewDeleteRollout(rolloutGateway ff_gateways.RolloutGateway) *DeleteRollout {
	return &DeleteRollout{rolloutGateway: rolloutGateway}
}

func (this *DeleteRollout) Execute(key string,
) ff_domains_exceptions.LibException {
	return this.rolloutGateway.Delete(key)
}
