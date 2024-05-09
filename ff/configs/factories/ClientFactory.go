/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_configs_factories

import (
	"fmt"
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways_beans "github.com/GabrielEstmr/ff-4-go/ff/gateways/beans"
	ff_gateways_ws_beans "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/beans"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
	ff_usecases_externalinterfaces_impl "github.com/GabrielEstmr/ff-4-go/ff/usecases/externalinterfaces/impl"
	ff_utils "github.com/GabrielEstmr/ff-4-go/ff/utils"
)

// ClientFactory is the factory of FfClient
type ClientFactory struct {
	_MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY string
	_MSG_ERROR_REGISTER_FEATURES_FLAGS               string
	_MSG_ERROR_REGISTER_FEATURES_PROPERTIES          string
	_MSG_ERROR_REGISTER_ROLLOUTS                     string
	errorUtils                                       ff_utils.ErrorUtils
	clientArgs                                       ff_configs_resources.FfClientArgs
}

// NewClientFactory is the constructor of ClientFactory
// By using this factory it is possible to generate a FfClient to use ff-4-go in your application and have access to:
//   - methods to manipulate Feature-Flags, Feature-Properties and Rollouts
//   - functions to Use in your Golang server
func NewClientFactory(clientArgs ff_configs_resources.FfClientArgs) *ClientFactory {
	return &ClientFactory{
		_MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY: "Error to get features register provider",
		_MSG_ERROR_REGISTER_FEATURES_FLAGS:               "error to register feature",
		_MSG_ERROR_REGISTER_FEATURES_PROPERTIES:          "error to register feature",
		_MSG_ERROR_REGISTER_ROLLOUTS:                     "error to register feature",
		errorUtils:                                       *new(ff_utils.ErrorUtils),
		clientArgs:                                       clientArgs,
	}
}

// Build builds the FfClient based on FfClientArgs
// An error in return indicates that the client could not be created or the initial state of the features or rollout could not be registered
func (this *ClientFactory) Build() (*ff_configs_resources.FfClient, error) {
	featureFlagGatewayImpl, errFm := ff_gateways_beans.NewFeatureFlagGatewayFactory(
		this.clientArgs).Get()
	if errFm != nil {
		fmt.Println(fmt.Sprintf(this._MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY, errFm.Error()))
		return new(ff_configs_resources.FfClient), errFm
	}

	featurePropertyGatewayImpl, errFp := ff_gateways_beans.NewFeaturePropertyGatewayFactory(
		this.clientArgs).Get()
	if errFp != nil {
		fmt.Println(fmt.Sprintf(this._MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY, errFp.Error()))
		return new(ff_configs_resources.FfClient), errFp
	}

	rolloutGatewayImpl, errFm := ff_gateways_beans.NewRolloutGatewayFactory(this.clientArgs).Get()
	if errFm != nil {
		fmt.Println(fmt.Sprintf(this._MSG_ERROR_INSTANTIATE_FEATURES_REGISTER_FACTORY, errFm.Error()))
		return new(ff_configs_resources.FfClient), errFm
	}

	useCaseBeans := *ff_usecases_beans.NewUseCaseBeans(
		featureFlagGatewayImpl,
		featurePropertyGatewayImpl,
		rolloutGatewayImpl)
	featureFlagMethods := ff_usecases_externalinterfaces_impl.NewFeatureFlagMethodsImpl(useCaseBeans)
	featurePropertyMethods := ff_usecases_externalinterfaces_impl.NewFeaturePropertyMethodsImpl(useCaseBeans)
	rolloutMethods := ff_usecases_externalinterfaces_impl.NewRolloutMethodsImpl(useCaseBeans)

	registerMethods := ff_usecases_externalinterfaces_impl.NewRegisterMethodsImpl(useCaseBeans)

	if this.clientArgs.HasInitialFeatureFlagValues() {
		errR := registerMethods.RegisterFeatureFlags(this.clientArgs.GetFeatureFlags())
		if errR != nil {
			fmt.Println(fmt.Sprintf(this._MSG_ERROR_REGISTER_FEATURES_FLAGS, errR.Error()))
			return new(ff_configs_resources.FfClient), errR
		}
	}

	if this.clientArgs.HasInitialFeaturePropertiesValues() {
		errR := registerMethods.RegisterFeatureProperties(this.clientArgs.GetFeatureProperties())
		if errR != nil {
			fmt.Println(fmt.Sprintf(this._MSG_ERROR_REGISTER_FEATURES_PROPERTIES, errR.Error()))
			return new(ff_configs_resources.FfClient), errR
		}
	}

	if this.clientArgs.HasInitialRolloutsValues() {
		errR := registerMethods.RegisterRollouts(this.clientArgs.GetRollouts())
		if errR != nil {
			fmt.Println(fmt.Sprintf(this._MSG_ERROR_REGISTER_ROLLOUTS, errR.Error()))
			return new(ff_configs_resources.FfClient), errR
		}
	}

	routeFns := *NewRouterFnsFactory(
		*(ff_gateways_ws_beans.NewControllerBeans(useCaseBeans)),
		this.clientArgs,
	)

	return ff_configs_resources.NewFfClient(
		&this.clientArgs,
		featureFlagMethods,
		featurePropertyMethods,
		rolloutMethods,
		routeFns.GetFunctionBeans()), nil
}
