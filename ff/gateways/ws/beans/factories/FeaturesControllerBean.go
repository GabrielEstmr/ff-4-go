package ff_gateways_ws_beans_factories

import (
	ff_configs_resources "baseapplicationgo/main/configs/ff/lib/configs/resources"
	"baseapplicationgo/main/configs/ff/lib/gateways/ws/controllers"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type FeaturesControllerBean struct {
	clientArgs   ff_configs_resources.FfClientArgs
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeaturesControllerBean(
	clientArgs ff_configs_resources.FfClientArgs,
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *FeaturesControllerBean {
	return &FeaturesControllerBean{
		clientArgs:   clientArgs,
		useCaseBeans: useCaseBeans,
	}
}

func (this FeaturesControllerBean) Get() *ff_gateways_ws_controllers.FeaturesController {
	return ff_gateways_ws_controllers.NewFeaturesController(
		this.clientArgs,
		this.useCaseBeans.CreateFeature,
		this.useCaseBeans.DeleteFeature,
		this.useCaseBeans.DisableFeature,
		this.useCaseBeans.EnableFeature,
		this.useCaseBeans.FindFeatureByKey,
		this.useCaseBeans.IsFeatureEnabled,
	)

}
