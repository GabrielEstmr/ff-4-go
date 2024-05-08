package ff_gateways_redis_repo

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways_redis_documents "baseapplicationgo/main/configs/ff/lib/gateways/redis/documents"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type FeaturePropertyRedisRepo struct {
	ffConfigData *ff.FfClientArgs
}

func NewFeaturePropertyRedisRepo(ffConfigData ff.FfClientArgs) *FeaturePropertyRedisRepo {
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
