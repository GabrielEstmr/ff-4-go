package ff_redis_repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_redis_documents "github.com/GabrielEstmr/ff-4-go/ff/redis/documents"
	"github.com/redis/go-redis/v9"
	"time"
)

type FeaturesRedisRepo struct {
	ffConfigData *ff.FfConfigData
}

func NewFeaturesRedisRepo(ffConfigData *ff.FfConfigData) *FeaturesRedisRepo {
	return &FeaturesRedisRepo{ffConfigData}
}

func (this *FeaturesRedisRepo) Save(
	feature ff_redis_documents.FeaturesDataRedisDocument) (
	ff_redis_documents.FeaturesDataRedisDocument, error) {

	featureBytes, err := json.Marshal(feature)
	if err != nil {
		return *new(ff_redis_documents.FeaturesDataRedisDocument), err
	}

	this.ffConfigData.GetCacheClient().Set(context.TODO(), this.buildKeyPrefix(feature.Key), featureBytes, time.Hour).Result()
	return feature, nil
}

func (this *FeaturesRedisRepo) FindById(key string) (
	ff_redis_documents.FeaturesDataRedisDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.buildKeyPrefix(key)).Result()

	if errors.Is(err, redis.Nil) {
		return *new(ff_redis_documents.FeaturesDataRedisDocument), nil
	}

	if err != nil {
		return *new(ff_redis_documents.FeaturesDataRedisDocument), err
	}

	var feature ff_redis_documents.FeaturesDataRedisDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(ff_redis_documents.FeaturesDataRedisDocument), err
	}

	return feature, nil
}

func (this *FeaturesRedisRepo) buildKeyPrefix(key string) string {
	return this.ffConfigData.GetCachingPrefix() + "_" + key
}
