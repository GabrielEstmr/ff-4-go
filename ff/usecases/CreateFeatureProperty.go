package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type CreateFeatureProperty struct {
	featurePropertyGateway ff_gateways.FeaturePropertyGateway
}

func NewCreateFeatureProperty(
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
) *CreateFeatureProperty {
	return &CreateFeatureProperty{
		featurePropertyGateway: featurePropertyGateway,
	}
}

func (this *CreateFeatureProperty) Execute(property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	persistedProperty, errF := this.featurePropertyGateway.FindById(property.GetKey())
	if errF != nil {
		return *new(ff_domains.FeatureProperty), errF
	}

	if !persistedProperty.IsEmpty() {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewConflictExceptionSglMsg(
				fmt.Sprintf("Feature property with key %s already exists.", property.GetKey()))
	}

	return this.featurePropertyGateway.Save(property)
}
