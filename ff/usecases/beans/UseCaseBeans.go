package ff_usecases_beans

import (
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_usecases "baseapplicationgo/main/configs/ff/lib/usecases"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
)

type UseCaseBeans struct {
	CreateFeature    ff_usecases_interfaces.CreateFeature
	DeleteFeature    ff_usecases_interfaces.DeleteFeature
	DisableFeature   ff_usecases_interfaces.DisableFeature
	EnableFeature    ff_usecases_interfaces.EnableFeature
	FindFeatureByKey ff_usecases_interfaces.FindFeatureByKey
	IsFeatureEnabled ff_usecases_interfaces.IsFeatureEnabled
	RegisterFeatures ff_usecases_interfaces.RegisterFeatures
}

func NewUseCaseBeans(
	featureGateway ff_gateways.FeaturesGateway,
) *UseCaseBeans {
	return &UseCaseBeans{
		CreateFeature:    ff_usecases.NewCreateFeature(featureGateway),
		DeleteFeature:    ff_usecases.NewDeleteFeature(featureGateway),
		DisableFeature:   ff_usecases.NewDisableFeature(featureGateway, ff_usecases.NewFindFeatureByKey(featureGateway)),
		EnableFeature:    ff_usecases.NewEnableFeature(featureGateway, ff_usecases.NewFindFeatureByKey(featureGateway)),
		FindFeatureByKey: ff_usecases.NewFindFeatureByKey(featureGateway),
		IsFeatureEnabled: ff_usecases.NewIsFeatureEnabled(ff_usecases.NewFindFeatureByKey(featureGateway)),
		RegisterFeatures: ff_usecases.NewRegisterFeatures(featureGateway),
	}
}
