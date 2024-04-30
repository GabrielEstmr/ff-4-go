package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type CreateFeature struct {
	featureGateway ff_gateways.FeaturesGateway
}

func NewCreateFeature(featureGateway ff_gateways.FeaturesGateway) *CreateFeature {
	return &CreateFeature{featureGateway: featureGateway}
}

func (this *CreateFeature) Execute(feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featureGateway.FindById(feature.GetKey())
	if errF != nil {
		return *new(ff_domains.Feature), errF
	}

	if !persistedFeature.IsEmpty() {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("Feature with key %s already exists.", feature.GetKey()))
	}

	return this.featureGateway.Save(feature)
}
