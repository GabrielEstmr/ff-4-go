package ff_configs_resources

type RouteFn map[string]Route

func (this RouteFn) GetRouteByUri(key string) Route {
	return this[key]
}

//func NewRouteFn(routes map[string]Route) *RouteFn {
//	return &RouteFn{routes: routes}
//}
//
//func (this RouteFn) GetRoutes() map[string]Route {
//	return this.routes
//}
