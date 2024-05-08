package ff_configs_resources

import (
	ff_usecases_externalinterfaces "baseapplicationgo/main/configs/ff/lib/usecases/externalinterfaces"
)

// FfClient is the client of ff-4-go in your application and allows to have access to:
//   - methods to manipulate Feature-Flags, Feature-Properties and Rollouts
//   - functions to Use in your Golang server
type FfClient struct {
	clientArgs              *FfClientArgs
	featuresMethods         ff_usecases_externalinterfaces.FeaturesMethods
	featuresPropertyMethods ff_usecases_externalinterfaces.FeaturePropertyMethods
	rolloutMethods          ff_usecases_externalinterfaces.RolloutMethods
	routeFn                 RouteFn
}

func NewFfClient(
	clientArgs *FfClientArgs,
	featuresMethods ff_usecases_externalinterfaces.FeaturesMethods,
	featuresPropertyMethods ff_usecases_externalinterfaces.FeaturePropertyMethods,
	rolloutMethods ff_usecases_externalinterfaces.RolloutMethods,
	routeFn RouteFn,
) *FfClient {
	return &FfClient{
		clientArgs:              clientArgs,
		featuresMethods:         featuresMethods,
		featuresPropertyMethods: featuresPropertyMethods,
		rolloutMethods:          rolloutMethods,
		routeFn:                 routeFn,
	}
}

func (this FfClient) GetClientArgs() *FfClientArgs {
	return this.clientArgs
}

func (this FfClient) GetFeaturesMethods() ff_usecases_externalinterfaces.FeaturesMethods {
	return this.featuresMethods
}

func (this FfClient) GetFeaturesPropertyMethods() ff_usecases_externalinterfaces.FeaturePropertyMethods {
	return this.featuresPropertyMethods
}

func (this FfClient) GetRolloutMethods() ff_usecases_externalinterfaces.RolloutMethods {
	return this.rolloutMethods
}

func (this FfClient) GetRouteFn() RouteFn {
	return this.routeFn
}
