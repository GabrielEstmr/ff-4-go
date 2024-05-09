/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_redis_repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways_redis_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis/documents"
	"github.com/redis/go-redis/v9"
	"time"
)

type FeaturePropertyRedisRepo struct {
	ffConfigData *ff_configs_resources.FfClientArgs
}

func NewFeaturePropertyRedisRepo(ffConfigData ff_configs_resources.FfClientArgs) *FeaturePropertyRedisRepo {
	return &FeaturePropertyRedisRepo{&ffConfigData}
}

func (this *FeaturePropertyRedisRepo) Save(
	propertyDocument ff_gateways_redis_documents.FeaturePropertyDocument) (
	ff_gateways_redis_documents.FeaturePropertyDocument, error) {

	featureBytes, err := json.Marshal(propertyDocument)
	if err != nil {
		return *new(ff_gateways_redis_documents.FeaturePropertyDocument), err
	}

	_, errS := this.ffConfigData.GetCacheClient().Set(
		context.TODO(), this.buildKeyPrefix(propertyDocument.Key), featureBytes, time.Hour).Result()
	if errS != nil {
		return *new(ff_gateways_redis_documents.FeaturePropertyDocument), errS
	}
	return propertyDocument, nil
}

func (this *FeaturePropertyRedisRepo) Delete(key string) error {
	del := this.ffConfigData.GetCacheClient().Del(context.TODO(), this.buildKeyPrefix(key))
	return del.Err()
}

func (this *FeaturePropertyRedisRepo) FindById(key string) (
	ff_gateways_redis_documents.FeaturePropertyDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.buildKeyPrefix(key)).Result()

	if errors.Is(err, redis.Nil) {
		return *new(ff_gateways_redis_documents.FeaturePropertyDocument), nil
	}

	if err != nil {
		return *new(ff_gateways_redis_documents.FeaturePropertyDocument), err
	}

	var feature ff_gateways_redis_documents.FeaturePropertyDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(ff_gateways_redis_documents.FeaturePropertyDocument), err
	}

	return feature, nil
}

func (this *FeaturePropertyRedisRepo) buildKeyPrefix(key string) string {
	return fmt.Sprintf("%s_%s_%s", this.ffConfigData.GetCachingPrefix(), this.ffConfigData.GetFeaturesColName(), key)
}
