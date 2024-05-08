package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type FindFeaturePropertyById struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewFindFeaturePropertyById(featurePropertyGateway ff_gateways.FeaturePropertyGateway) *FindFeaturePropertyById {
	return &FindFeaturePropertyById{featurePropertyGateway: featurePropertyGateway}
}

func (this *FindFeaturePropertyById) Execute(key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featurePropertyGateway.FindById(key)
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	if persistedFeature.IsEmpty() {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("feature property with key %s not found.", key))
	}

	return persistedFeature, nil
}
