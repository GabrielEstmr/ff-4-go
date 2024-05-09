package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
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
