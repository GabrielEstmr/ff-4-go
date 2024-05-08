package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type UpdateFeatureProperty struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewUpdateFeatureProperty(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *UpdateFeatureProperty {
	return &UpdateFeatureProperty{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *UpdateFeatureProperty) Execute(property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedProperty, errF := this.featurePropertyGateway.FindById(property.GetKey())
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	if persistedProperty.IsEmpty() {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				fmt.Sprintf("feature property with key %s do not exists.", property.GetKey()))
	}

	return this.featurePropertyGateway.Update(property)
}
