/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_usecases_interfaces

import ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"

type DeleteRollout interface {
	Execute(key string,
	) ff_domains_exceptions.LibException
}
