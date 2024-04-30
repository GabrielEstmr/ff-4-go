package ff_gateways_ws_beans

import (
	ff_configs_resources "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways_ws_beans_factories "baseapplicationgo/main/configs/ff/lib/gateways/ws/beans/factories"
	"baseapplicationgo/main/configs/ff/lib/gateways/ws/controllers"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type ControllerBeans struct {
	FeaturesController *ff_gateways_ws_controllers.FeaturesController
}

func NewControllerBeans(
	clientArgs ff_configs_resources.FfClientArgs,
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *ControllerBeans {
	return &ControllerBeans{
		FeaturesController: ff_gateways_ws_beans_factories.NewFeaturesControllerBean(clientArgs, useCaseBeans).Get(),
	}
}
