package ff_gateways_ws_beans_factories

import (
	ff_gateways_ws_controllers "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/controllers"
	ff_usecases_beans "github.com/GabrielEstmr/ff-4-go/ff/usecases/beans"
)

type FeaturePropertyControllerBean struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewFeaturePropertyControllerBean(
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *FeaturePropertyControllerBean {
	return &FeaturePropertyControllerBean{
		useCaseBeans: useCaseBeans,
	}
}

func (this FeaturePropertyControllerBean) Get() *ff_gateways_ws_controllers.FeaturePropertyController {
	return ff_gateways_ws_controllers.NewFeaturePropertyController(
		this.useCaseBeans.CreateFeatureProperty,
		this.useCaseBeans.UpdateFeatureProperty,
		this.useCaseBeans.DeleteFeatureProperty,
		this.useCaseBeans.FindFeaturePropertyById,
		this.useCaseBeans.AddValueToFeatureProperty,
		this.useCaseBeans.RemoveValueToFeatureProperty,
		this.useCaseBeans.EnableFeatureProperty,
		this.useCaseBeans.DisableFeatureProperty,
	)
}
