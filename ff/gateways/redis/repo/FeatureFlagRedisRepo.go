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

type FeatureFlagRedisRepo struct {
	ffConfigData *ff.FfClientArgs
}

func NewFeatureFlagRedisRepo(ffConfigData ff.FfClientArgs) *FeatureFlagRedisRepo {
	return &FeatureFlagRedisRepo{&ffConfigData}
}

func (this *FeatureFlagRedisRepo) Save(
	featureFlagDocument ff_gateways_redis_documents.FeatureFlagDocument) (
	ff_gateways_redis_documents.FeatureFlagDocument, error) {

	featureBytes, err := json.Marshal(featureFlagDocument)
	if err != nil {
		return *new(ff_gateways_redis_documents.FeatureFlagDocument), err
	}

	_, errS := this.ffConfigData.GetCacheClient().Set(
		context.TODO(), this.buildKeyPrefix(featureFlagDocument.Key), featureBytes, time.Hour).Result()
	if errS != nil {
		return *new(ff_gateways_redis_documents.FeatureFlagDocument), err
	}
	return featureFlagDocument, nil
}

func (this *FeatureFlagRedisRepo) Delete(key string) error {
	del := this.ffConfigData.GetCacheClient().Del(context.TODO(), this.buildKeyPrefix(key))
	return del.Err()
}

func (this *FeatureFlagRedisRepo) FindById(key string) (
	ff_gateways_redis_documents.FeatureFlagDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.buildKeyPrefix(key)).Result()

	if errors.Is(err, redis.Nil) {
		return *new(ff_gateways_redis_documents.FeatureFlagDocument), nil
	}

	if err != nil {
		return *new(ff_gateways_redis_documents.FeatureFlagDocument), err
	}

	var feature ff_gateways_redis_documents.FeatureFlagDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(ff_gateways_redis_documents.FeatureFlagDocument), err
	}

	return feature, nil
}

func (this *FeatureFlagRedisRepo) buildKeyPrefix(key string) string {
	return fmt.Sprintf("%s_%s_%s", this.ffConfigData.GetCachingPrefix(), this.ffConfigData.GetFeaturesFlagColName(), key)
}