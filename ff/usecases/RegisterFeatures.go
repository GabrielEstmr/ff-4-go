package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type RegisterFeatures struct {
	featureGateway ff_gateways.FeaturesGateway
}

func NewRegisterFeatures(featureGateway ff_gateways.FeaturesGateway) *RegisterFeatures {
	return &RegisterFeatures{featureGateway: featureGateway}
}

func (this *RegisterFeatures) Execute(features ff_domains.Features) ff_domains_exceptions.LibException {
	for k, v := range features {
		persistedFeature, errF := this.featureGateway.FindById(v.GetKey())
		if errF != nil {
			fmt.Printf("error to find key of name %s\n", k)
			return errF
		}
		if persistedFeature.IsEmpty() {
			_, errS := this.featureGateway.Save(v)
			if errS != nil {
				fmt.Printf("error to save key of name %s\n", k)
				return errS
			}
			fmt.Printf("feature with key %s has been saved", v.GetKey())
		}
	}
	return nil
}
