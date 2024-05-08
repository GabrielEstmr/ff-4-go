package ff_gateways_ws_beans_factories

import (
	ff_gateways_ws_controllers "baseapplicationgo/main/configs/ff/lib/gateways/ws/controllers"
	ff_usecases_beans "baseapplicationgo/main/configs/ff/lib/usecases/beans"
)

type RolloutControllerBean struct {
	useCaseBeans ff_usecases_beans.UseCaseBeans
}

func NewRolloutControllerBean(
	useCaseBeans ff_usecases_beans.UseCaseBeans,
) *RolloutControllerBean {
	return &RolloutControllerBean{
		useCaseBeans: useCaseBeans,
	}
}

func (this RolloutControllerBean) Get() *ff_gateways_ws_controllers.RolloutController {
	return ff_gateways_ws_controllers.NewRolloutController(
		this.useCaseBeans.CreateRollout,
		this.useCaseBeans.UpdateRollout,
		this.useCaseBeans.DeleteRollout,
		this.useCaseBeans.FindRolloutById,
		this.useCaseBeans.AddTargetToRollout,
		this.useCaseBeans.RemoveTargetFromRollout,
		this.useCaseBeans.EnableRolloutToAll,
		this.useCaseBeans.DisableRolloutToAll,
		this.useCaseBeans.VerifyIsEnabledAllOrTargetInRollout,
	)
}
