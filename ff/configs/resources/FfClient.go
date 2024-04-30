package ff_configs_resources

import (
	ff_usecases_externalinterfaces "baseapplicationgo/main/configs/ff/lib/usecases/externalinterfaces"
)

type FfClient struct {
	clientArgs      *FfClientArgs
	featuresMethods ff_usecases_externalinterfaces.FeaturesMethods
	registerMethods ff_usecases_externalinterfaces.RegisterMethods
	routeFn         RouteFn
}

func NewFfClient(
	clientArgs *FfClientArgs,
	featuresMethods ff_usecases_externalinterfaces.FeaturesMethods,
	registerMethods ff_usecases_externalinterfaces.RegisterMethods,
	routeFn RouteFn,
) *FfClient {
	return &FfClient{
		clientArgs:      clientArgs,
		featuresMethods: featuresMethods,
		registerMethods: registerMethods,
		routeFn:         routeFn,
	}
}

func (this FfClient) GetClientArgs() *FfClientArgs {
	return this.clientArgs
}

func (this FfClient) GetFeaturesMethods() ff_usecases_externalinterfaces.FeaturesMethods {
	return this.featuresMethods
}

func (this FfClient) GetRegisterMethods() ff_usecases_externalinterfaces.RegisterMethods {
	return this.registerMethods
}

func (this FfClient) GetRouteFn() RouteFn {
	return this.routeFn
}
