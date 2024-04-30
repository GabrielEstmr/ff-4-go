package ff_configs_factories

import (
	ff_configs_resources "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways_ws_beans "baseapplicationgo/main/configs/ff/lib/gateways/ws/beans"
	"net/http"
)

type RouterFnsFactory struct {
	controllerBeans ff_gateways_ws_beans.ControllerBeans
	clientArgs      ff_configs_resources.FfClientArgs
}

func NewRouterFnsFactory(
	controllerBeans ff_gateways_ws_beans.ControllerBeans,
	clientArgs ff_configs_resources.FfClientArgs,
) *RouterFnsFactory {
	return &RouterFnsFactory{
		controllerBeans: controllerBeans,
		clientArgs:      clientArgs,
	}
}

func (this RouterFnsFactory) GetFunctionBeans() ff_configs_resources.RouteFn {
	values := make(map[string]ff_configs_resources.Route)
	this.appendRoute(
		values,
		"/v1/features/{key}",
		http.MethodPost,
		this.controllerBeans.FeaturesController.CreateFeature,
	)
	this.appendRoute(
		values,
		"/v1/features/{key}",
		http.MethodDelete,
		this.controllerBeans.FeaturesController.DeleteFeature,
	)
	this.appendRoute(
		values,
		"/v1/features/{key}/disable",
		http.MethodPost,
		this.controllerBeans.FeaturesController.DisableFeature,
	)
	this.appendRoute(
		values,
		"/v1/features/{key}/enable",
		http.MethodPost,
		this.controllerBeans.FeaturesController.EnableFeature,
	)
	this.appendRoute(
		values,
		"/v1/features/{key}",
		http.MethodGet,
		this.controllerBeans.FeaturesController.FindFeatureByKey,
	)
	this.appendRoute(
		values,
		"/v1/features/{key}/verify-enabled",
		http.MethodPost,
		this.controllerBeans.FeaturesController.IsFeatureEnabled,
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
