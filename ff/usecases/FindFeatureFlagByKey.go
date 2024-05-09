package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
)

type FindFeatureFlagByKey struct {
	featureGateway ff_gateways.FeatureFlagsGateway
}

func NewFindFeatureFlagByKey(featureGateway ff_gateways.FeatureFlagsGateway) *FindFeatureFlagByKey {
	return &FindFeatureFlagByKey{featureGateway: featureGateway}
}

func (this *FindFeatureFlagByKey) Execute(key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featureGateway.FindById(key)
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	if persistedFeature.IsEmpty() {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("feature with key %s not found.", key))
	}

	return persistedFeature, nil
}
