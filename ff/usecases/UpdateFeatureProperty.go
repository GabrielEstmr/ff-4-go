package ff_usecases

import (
	"fmt"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
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
