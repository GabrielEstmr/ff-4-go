package ff_configs_factories

import (
	ff_configs_resources "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_factories "baseapplicationgo/main/configs/ff/lib/factories"
	ff_gateways_ws_beans "baseapplicationgo/main/configs/ff/lib/gateways/ws/beans"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
	ff_usecases_externalinterfaces_impl "baseapplicationgo/main/configs/ff/lib/usecases/externalinterfaces/impl"
	ff_utils "baseapplicationgo/main/configs/ff/lib/utils"
)

type ClientFactory struct {
	_MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY string
	_MSG_ERROR_REGISTER_FEATURES                     string
	errorUtils                                       ff_utils.ErrorUtils
	clientArgs                                       ff_configs_resources.FfClientArgs
}

func NewClientFactory(clientArgs ff_configs_resources.FfClientArgs) *ClientFactory {
	return &ClientFactory{
		_MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY: "Error to get features register provider",
		_MSG_ERROR_REGISTER_FEATURES:                     "error to register feature",
		errorUtils:                                       *new(ff_utils.ErrorUtils),
		clientArgs:                                       clientArgs,
	}
}

func (this *ClientFactory) Build() *ff_configs_resources.FfClient {
	featureMethodsImpl, errFm := ff_factories.NewFeaturesGatewayFactory(this.clientArgs).Get()
	this.errorUtils.FailOnError(errFm, this._MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY)

	useCaseBeans := *ff_usecases_beans.NewUseCaseBeans(featureMethodsImpl)
	featureMethods := ff_usecases_externalinterfaces_impl.NewFeaturesMethodsImpl(useCaseBeans)
	registerMethods := ff_usecases_externalinterfaces_impl.NewRegisterMethodsImpl(useCaseBeans)

	if this.clientArgs.HasInitialValues() {
		errR := registerMethods.RegisterFeature(this.clientArgs.GetFeaturesInitialValues())
		this.errorUtils.FailOnError(errR, this._MSG_ERROR_REGISTER_FEATURES)
	}

	routeFns := *NewRouterFnsFactory(
		*(ff_gateways_ws_beans.NewControllerBeans(this.clientArgs, useCaseBeans)),
		this.clientArgs,
	)

	return ff_configs_resources.NewFfClient(
		&this.clientArgs,
		featureMethods,
		registerMethods,
		routeFns.GetFunctionBeans())
}
