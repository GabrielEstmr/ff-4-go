package ff_configs_resources

type RouteFn map[string]Route

func (this RouteFn) GetRouteByUri(key string) Route {
	return this[key]
}
