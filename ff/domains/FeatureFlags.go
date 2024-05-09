/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_domains

// FeatureFlags are a map of FeatureFlag, where the keys of its elements are given by the key of each FeatureFlag.
type FeatureFlags map[string]FeatureFlag
