package ff_configs_factories

import (
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways_ws_beans "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/beans"
	"net/http"
)

type RouterFnsFactory struct {
	_FEATURE_FLAG_RESOURCE_NAME string
	_ROLLOUT_RESOURCE_NAME      string
	_FEATURE_PROP_RESOURCE_NAME string
	controllerBeans             ff_gateways_ws_beans.ControllerBeans
	clientArgs                  ff_configs_resources.FfClientArgs
}

func NewRouterFnsFactory(
	controllerBeans ff_gateways_ws_beans.ControllerBeans,
	clientArgs ff_configs_resources.FfClientArgs,
) *RouterFnsFactory {
	return &RouterFnsFactory{
		_FEATURE_FLAG_RESOURCE_NAME: "feature-flags",
		_ROLLOUT_RESOURCE_NAME:      "rollouts",
		_FEATURE_PROP_RESOURCE_NAME: "feature-properties",
		controllerBeans:             controllerBeans,
		clientArgs:                  clientArgs,
	}
}

func (this RouterFnsFactory) GetFunctionBeans() ff_configs_resources.RouteFn {
	values := make(map[string]ff_configs_resources.Route)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME,
		http.MethodPost,
		this.controllerBeans.FeaturesController.CreateFeature,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}",
		http.MethodDelete,
		this.controllerBeans.FeaturesController.DeleteFeature,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}/disable",
		http.MethodPost,
		this.controllerBeans.FeaturesController.DisableFeature,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}/enable",
		http.MethodPost,
		this.controllerBeans.FeaturesController.EnableFeature,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}",
		http.MethodGet,
		this.controllerBeans.FeaturesController.FindFeatureByKey,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}/verify-enabled",
		http.MethodPost,
		this.controllerBeans.FeaturesController.IsFeatureEnabled,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_FLAG_RESOURCE_NAME+"/{key}/verify-disabled",
		http.MethodPost,
		this.controllerBeans.FeaturesController.IsFeatureDisabled,
	)

	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME,
		http.MethodPost,
		this.controllerBeans.RolloutController.CreateRollout,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}",
		http.MethodPut,
		this.controllerBeans.RolloutController.UpdateRollout,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}",
		http.MethodDelete,
		this.controllerBeans.RolloutController.DeleteRollout,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}",
		http.MethodGet,
		this.controllerBeans.RolloutController.FindRolloutByKey,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}/targets/{target}/add",
		http.MethodPut,
		this.controllerBeans.RolloutController.AddTargetToRollout,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}/targets/{target}/remove",
		http.MethodPut,
		this.controllerBeans.RolloutController.RemoveTargetFromRollout,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}/enable",
		http.MethodPost,
		this.controllerBeans.RolloutController.EnableToAll,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}/disable",
		http.MethodPost,
		this.controllerBeans.RolloutController.DisableToAll,
	)
	this.appendRoute(
		values,
		"/v1/"+this._ROLLOUT_RESOURCE_NAME+"/{key}/targets/{target}/verify",
		http.MethodPost,
		this.controllerBeans.RolloutController.VerifyIsTargetInRollout,
	)

	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME,
		http.MethodPost,
		this.controllerBeans.FeaturePropertyController.CreateFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}",
		http.MethodPut,
		this.controllerBeans.FeaturePropertyController.UpdateFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}",
		http.MethodDelete,
		this.controllerBeans.FeaturePropertyController.DeleteFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}",
		http.MethodGet,
		this.controllerBeans.FeaturePropertyController.FindFeaturePropertyById,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}/values/{value}/add",
		http.MethodPut,
		this.controllerBeans.FeaturePropertyController.AddValueToFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}/values/{value}/remove",
		http.MethodPut,
		this.controllerBeans.FeaturePropertyController.RemoveValueToFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}/enable",
		http.MethodPost,
		this.controllerBeans.FeaturePropertyController.EnableFeatureProperty,
	)
	this.appendRoute(
		values,
		"/v1/"+this._FEATURE_PROP_RESOURCE_NAME+"/{key}/disable",
		http.MethodPost,
		this.controllerBeans.FeaturePropertyController.DisableFeatureProperty,
	)
	return values
}

func (this RouterFnsFactory) buildUri(uriPath string) string {
	return this.clientArgs.GetBaseRoutePath() + uriPath
}

func (this RouterFnsFactory) appendRoute(
	values map[string]ff_configs_resources.Route,
	uriPath string,
	method string,
	fn func(w http.ResponseWriter, r *http.Request),
) {
	values[method+"_"+this.buildUri(uriPath)] = this.buildRoute(uriPath, method, fn)
}

func (this RouterFnsFactory) buildRoute(
	uriPath string,
	method string,
	fn func(w http.ResponseWriter, r *http.Request),
) ff_configs_resources.Route {
	return ff_configs_resources.Route{
		URI:            this.buildUri(uriPath),
		Method:         method,
		ControllerFunc: fn,
	}
}
