/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_configs_resources

type CacheType string

const (
	REDIS CacheType = "REDIS"
)

var CacheTypeEnum = map[CacheType]CacheType{
	REDIS: REDIS,
}

var CacheTypeEnumFromNames = map[string]CacheType{
	"REDIS": REDIS,
}

func (this CacheType) Exists() bool {
	_, exists := CacheTypeEnum[this]
	return exists
}

func (this CacheType) FromValue(value string) CacheType {
	valueMap, exists := CacheTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this CacheType) Name() string {
	valueMap, exists := CacheTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}
