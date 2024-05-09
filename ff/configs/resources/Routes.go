/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_configs_resources

type RouteFn map[string]Route

func (this RouteFn) GetRouteByUri(key string) Route {
	return this[key]
}
