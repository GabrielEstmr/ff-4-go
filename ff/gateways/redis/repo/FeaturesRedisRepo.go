package ff_gateways_redis_repo

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways_redis_documents "baseapplicationgo/main/configs/ff/lib/gateways/redis/documents"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type FeaturesRedisRepo struct {
	ffConfigData *ff.FfClientArgs
}

func NewFeaturesRedisRepo(ffConfigData ff.FfClientArgs) *FeaturesRedisRepo {
	return &FeaturesRedisRepo{&ffConfigData}
}

func (this *FeaturesRedisRepo) Save(
	feature ff_gateways_redis_documents.FeaturesDataRedisDocument) (
	ff_gateways_redis_documents.FeaturesDataRedisDocument, error) {

	featureBytes, err := json.Marshal(feature)
	if err != nil {
		return *new(ff_gateways_redis_documents.FeaturesDataRedisDocument), err
	}

	this.ffConfigData.GetCacheClient().Set(
		context.TODO(), this.buildKeyPrefix(feature.Key), featureBytes, time.Hour).Result()
	return feature, nil
}

func (this *FeaturesRedisRepo) Delete(key string) error {
	del := this.ffConfigData.GetCacheClient().Del(context.TODO(), this.buildKeyPrefix(key))
	return del.Err()
}

func (this *FeaturesRedisRepo) FindById(key string) (
	ff_gateways_redis_documents.FeaturesDataRedisDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.buildKeyPrefix(key)).Result()

	if errors.Is(err, redis.Nil) {
		return *new(ff_gateways_redis_documents.FeaturesDataRedisDocument), nil
	}

	if err != nil {
		return *new(ff_gateways_redis_documents.FeaturesDataRedisDocument), err
	}

	var feature ff_gateways_redis_documents.FeaturesDataRedisDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(ff_gateways_redis_documents.FeaturesDataRedisDocument), err
	}

	return feature, nil
}

func (this *FeaturesRedisRepo) buildKeyPrefix(key string) string {
	return this.ffConfigData.GetCachingPrefix() + "_" + key
}
