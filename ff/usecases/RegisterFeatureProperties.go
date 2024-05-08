package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type RegisterFeatureProperties struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewRegisterFeatureProperties(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *RegisterFeatureProperties {
	return &RegisterFeatureProperties{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *RegisterFeatureProperties) Execute(
	featureProperties ff_domains.FeatureProperties) ff_domains_exceptions.LibException {
	if featureProperties == nil {
		return nil
	}
	for k, v := range featureProperties {
		persistedFeature, errF := this.featurePropertyGateway.FindById(v.GetKey())
		if errF != nil {
			fmt.Printf("error to find key of name %s\n", k)
			return errF
		}
		if persistedFeature.IsEmpty() {
			_, errS := this.featurePropertyGateway.Save(v)
			if errS != nil {
				fmt.Printf("error to save key of name %s\n", k)
				return errS
			}
			fmt.Printf("feature property with key %s has been saved", v.GetKey())
		}
	}
	return nil
}
