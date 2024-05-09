/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_beans

import (
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_usecases "github.com/GabrielEstmr/ff-4-go/ff/usecases"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
)

type UseCaseBeans struct {
	CreateFeatureFlag                   ff_usecases_interfaces.CreateFeatureFlag
	DeleteFeatureFlag                   ff_usecases_interfaces.DeleteFeatureFlag
	DisableFeatureFlag                  ff_usecases_interfaces.DisableFeatureFlag
	EnableFeatureFlag                   ff_usecases_interfaces.EnableFeatureFlag
	FindFeatureFlagByKey                ff_usecases_interfaces.FindFeatureFlagByKey
	IsFeatureFlagEnabled                ff_usecases_interfaces.IsFeatureFlagEnabled
	IsFeatureFlagDisabled               ff_usecases_interfaces.IsFeatureFlagDisabled
	RegisterFeatureFlags                ff_usecases_interfaces.RegisterFeatureFlags
	AddTargetToRollout                  ff_usecases_interfaces.AddTargetToRollout
	CreateRollout                       ff_usecases_interfaces.CreateRollout
	DeleteRollout                       ff_usecases_interfaces.DeleteRollout
	DisableRolloutToAll                 ff_usecases_interfaces.DisableRolloutToAll
	EnableRolloutToAll                  ff_usecases_interfaces.EnableRolloutToAll
	FindRolloutById                     ff_usecases_interfaces.FindRolloutById
	RemoveTargetFromRollout             ff_usecases_interfaces.RemoveTargetFromRollout
	UpdateRollout                       ff_usecases_interfaces.UpdateRollout
	VerifyIsEnabledAllOrTargetInRollout ff_usecases_interfaces.VerifyIsEnabledAllOrTargetInRollout
	CreateFeatureProperty               ff_usecases_interfaces.CreateFeatureProperty
	UpdateFeatureProperty               ff_usecases_interfaces.UpdateFeatureProperty
	DeleteFeatureProperty               ff_usecases_interfaces.DeleteFeatureProperty
	FindFeaturePropertyById             ff_usecases_interfaces.FindFeaturePropertyById
	AddValueToFeatureProperty           ff_usecases_interfaces.AddValueToFeatureProperty
	RemoveValueToFeatureProperty        ff_usecases_interfaces.RemoveValueToFeatureProperty
	EnableFeatureProperty               ff_usecases_interfaces.EnableFeatureProperty
	DisableFeatureProperty              ff_usecases_interfaces.DisableFeatureProperty
	RegisterFeatureProperties           ff_usecases_interfaces.RegisterFeatureProperties
	RegisterRollouts                    ff_usecases_interfaces.RegisterRollouts
}

func NewUseCaseBeans(
	featureGateway ff_gateways.FeatureFlagsGateway,
	featurePropertyGateway ff_gateways.FeaturePropertyGateway,
	rolloutGateway ff_gateways.RolloutGateway,
) *UseCaseBeans {
	return &UseCaseBeans{
		CreateFeatureFlag:                   ff_usecases.NewCreateFeatureFlag(featureGateway),
		DeleteFeatureFlag:                   ff_usecases.NewDeleteFeatureFlag(featureGateway),
		DisableFeatureFlag:                  ff_usecases.NewDisableFeatureFlag(featureGateway, ff_usecases.NewFindFeatureFlagByKey(featureGateway)),
		EnableFeatureFlag:                   ff_usecases.NewEnableFeatureFlag(featureGateway, ff_usecases.NewFindFeatureFlagByKey(featureGateway)),
		FindFeatureFlagByKey:                ff_usecases.NewFindFeatureFlagByKey(featureGateway),
		IsFeatureFlagEnabled:                ff_usecases.NewIsFeatureFlagEnabled(ff_usecases.NewFindFeatureFlagByKey(featureGateway)),
		IsFeatureFlagDisabled:               ff_usecases.NewIsFeatureFlagDisabled(ff_usecases.NewIsFeatureFlagEnabled(ff_usecases.NewFindFeatureFlagByKey(featureGateway))),
		RegisterFeatureFlags:                ff_usecases.NewRegisterFeatureFlags(featureGateway),
		AddTargetToRollout:                  ff_usecases.NewAddTargetToRollout(ff_usecases.NewFindRolloutById(rolloutGateway), rolloutGateway),
		CreateRollout:                       ff_usecases.NewCreateRollout(rolloutGateway),
		DeleteRollout:                       ff_usecases.NewDeleteRollout(rolloutGateway),
		DisableRolloutToAll:                 ff_usecases.NewDisableRolloutToAll(ff_usecases.NewFindRolloutById(rolloutGateway), rolloutGateway),
		EnableRolloutToAll:                  ff_usecases.NewEnableRolloutToAll(ff_usecases.NewFindRolloutById(rolloutGateway), rolloutGateway),
		FindRolloutById:                     ff_usecases.NewFindRolloutById(rolloutGateway),
		RemoveTargetFromRollout:             ff_usecases.NewRemoveTargetFromRollout(ff_usecases.NewFindRolloutById(rolloutGateway), rolloutGateway),
		UpdateRollout:                       ff_usecases.NewUpdateRollout(rolloutGateway),
		VerifyIsEnabledAllOrTargetInRollout: ff_usecases.NewVerifyIsEnabledAllOrTargetInRollout(ff_usecases.NewFindRolloutById(rolloutGateway)),
		CreateFeatureProperty:               ff_usecases.NewCreateFeatureProperty(featurePropertyGateway),
		UpdateFeatureProperty:               ff_usecases.NewUpdateFeatureProperty(featurePropertyGateway),
		DeleteFeatureProperty:               ff_usecases.NewDeleteFeatureProperty(featurePropertyGateway),
		FindFeaturePropertyById:             ff_usecases.NewFindFeaturePropertyById(featurePropertyGateway),
		AddValueToFeatureProperty:           ff_usecases.NewAddValueToFeatureProperty(ff_usecases.NewFindFeaturePropertyById(featurePropertyGateway), featurePropertyGateway),
		RemoveValueToFeatureProperty:        ff_usecases.NewRemoveValueToFeatureProperty(ff_usecases.NewFindFeaturePropertyById(featurePropertyGateway), featurePropertyGateway),
		EnableFeatureProperty:               ff_usecases.NewEnableFeatureProperty(ff_usecases.NewFindFeaturePropertyById(featurePropertyGateway), featurePropertyGateway),
		DisableFeatureProperty:              ff_usecases.NewDisableFeatureProperty(ff_usecases.NewFindFeaturePropertyById(featurePropertyGateway), featurePropertyGateway),
		RegisterFeatureProperties:           ff_usecases.NewRegisterFeatureProperties(featurePropertyGateway),
		RegisterRollouts:                    ff_usecases.NewRegisterRollouts(rolloutGateway),
	}
}
