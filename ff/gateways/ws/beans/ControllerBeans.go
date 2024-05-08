package ff_gateways_ws_beans

import (
	ff_gateways_ws_beans_factories "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/beans/factories"
	ff_gateways_ws_controllers "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/controllers"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
)

type ControllerBeans struct {
	FeaturesController        *ff_gateways_ws_controllers.FeatureController
	RolloutController         *ff_gateways_ws_controllers.RolloutController
	FeaturePropertyController *ff_gateways_ws_controllers.FeaturePropertyController
}

func NewControllerBeans(
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *ControllerBeans {
	return &ControllerBeans{
		FeaturesController:        ff_gateways_ws_beans_factories.NewFeatureControllerBean(useCaseBeans).Get(),
		RolloutController:         ff_gateways_ws_beans_factories.NewRolloutControllerBean(useCaseBeans).Get(),
		FeaturePropertyController: ff_gateways_ws_beans_factories.NewFeaturePropertyControllerBean(useCaseBeans).Get(),
	}
}
