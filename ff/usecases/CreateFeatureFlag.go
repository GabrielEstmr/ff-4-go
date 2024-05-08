package ff_usecases

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"fmt"
)

type CreateFeatureFlag struct {
	featureGateway ff_gateways.FeatureFlagsGateway
}

func NewCreateFeatureFlag(featureGateway ff_gateways.FeatureFlagsGateway) *CreateFeatureFlag {
	return &CreateFeatureFlag{featureGateway: featureGateway}
}

func (this *CreateFeatureFlag) Execute(featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	persistedFeature, errF := this.featureGateway.FindById(featureFlag.GetKey())
	if errF != nil {
		return *new(ff_domains.FeatureFlag), errF
	}

	if !persistedFeature.IsEmpty() {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewConflictExceptionSglMsg(
				fmt.Sprintf("Feature with key %s already exists.", featureFlag.GetKey()))
	}

	return this.featureGateway.Save(featureFlag)
}
