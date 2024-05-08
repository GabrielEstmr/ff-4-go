package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type RegisterRollouts struct {
	rolloutGateway ff_gateways.RolloutGateway
}

func NewRegisterRollouts(
	rolloutGateway ff_gateways.RolloutGateway,
) *RegisterRollouts {
	return &RegisterRollouts{
		rolloutGateway: rolloutGateway,
	}
}

func (this *RegisterRollouts) Execute(
	rollouts ff_domains.Rollouts) ff_domains_exceptions.LibException {
	if rollouts == nil {
		return nil
	}
	for k, v := range rollouts {
		persistedFeature, errF := this.rolloutGateway.FindById(v.GetKey())
		if errF != nil {
			fmt.Printf("error to find key of name %s\n", k)
			return errF
		}
		if persistedFeature.IsEmpty() {
			_, errS := this.rolloutGateway.Save(v)
			if errS != nil {
				fmt.Printf("error to save key of name %s\n", k)
				return errS
			}
			fmt.Printf("feature property with key %s has been saved", v.GetKey())
		}
	}
	return nil
}
