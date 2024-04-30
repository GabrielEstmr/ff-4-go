package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type FindFeatureByKey struct {
	featureGateway ff_gateways.FeaturesGateway
}

func NewFindFeatureByKey(featureGateway ff_gateways.FeaturesGateway) *FindFeatureByKey {
	return &FindFeatureByKey{featureGateway: featureGateway}
}

func (this *FindFeatureByKey) Execute(key string,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featureGateway.FindById(key)
	if errF != nil {
		return *new(ff_domains.Feature), errF
	}

	if persistedFeature.IsEmpty() {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("Feature with key %s not found.", key))
	}

	return persistedFeature, nil
}
