package ff_usecases

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
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
