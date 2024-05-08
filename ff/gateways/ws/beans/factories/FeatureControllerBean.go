package ff_gateways_ws_beans_factories

import (
	"baseapplicationgo/main/configs/ff/lib/gateways/ws/controllers"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type FeatureControllerBean struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeatureControllerBean(
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *FeatureControllerBean {
	return &FeatureControllerBean{
		useCaseBeans: useCaseBeans,
	}
}

func (this FeatureControllerBean) Get() *ff_gateways_ws_controllers.FeatureController {
	return ff_gateways_ws_controllers.NewFeatureController(
		this.useCaseBeans.CreateFeatureFlag,
		this.useCaseBeans.DeleteFeatureFlag,
		this.useCaseBeans.DisableFeatureFlag,
		this.useCaseBeans.EnableFeatureFlag,
		this.useCaseBeans.FindFeatureFlagByKey,
		this.useCaseBeans.IsFeatureFlagEnabled,
		this.useCaseBeans.IsFeatureFlagDisabled,
	)
}
