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

type RolloutRedisRepo struct {
	ffConfigData *ff.FfClientArgs
}

func NewRolloutRedisRepo(ffConfigData ff.FfClientArgs) *RolloutRedisRepo {
	return &RolloutRedisRepo{&ffConfigData}
}

func (this *RolloutRedisRepo) Save(
	rolloutDocument ff_gateways_redis_documents.RolloutDocument) (
	ff_gateways_redis_documents.RolloutDocument, error) {

	featureBytes, err := json.Marshal(rolloutDocument)
	if err != nil {
		return *new(ff_gateways_redis_documents.RolloutDocument), err
	}

	_, errS := this.ffConfigData.GetCacheClient().Set(
		context.TODO(), this.buildKeyPrefix(rolloutDocument.Key), featureBytes, time.Hour).Result()
	if errS != nil {
		return *new(ff_gateways_redis_documents.RolloutDocument), errS
	}
	return rolloutDocument, nil
}

func (this *RolloutRedisRepo) Delete(key string) error {
	del := this.ffConfigData.GetCacheClient().Del(context.TODO(), this.buildKeyPrefix(key))
	return del.Err()
}

func (this *RolloutRedisRepo) FindById(key string) (
	ff_gateways_redis_documents.RolloutDocument, error) {

	result, err := this.ffConfigData.GetCacheClient().
		Get(context.TODO(), this.buildKeyPrefix(key)).Result()

	if errors.Is(err, redis.Nil) {
		return *new(ff_gateways_redis_documents.RolloutDocument), nil
	}

	if err != nil {
		return *new(ff_gateways_redis_documents.RolloutDocument), err
	}

	var feature ff_gateways_redis_documents.RolloutDocument
	if err = json.Unmarshal([]byte(result), &feature); err != nil {
		return *new(ff_gateways_redis_documents.RolloutDocument), err
	}

	return feature, nil
}

func (this *RolloutRedisRepo) buildKeyPrefix(key string) string {
	return fmt.Sprintf("%s_%s_%s", this.ffConfigData.GetCachingPrefix(), this.ffConfigData.GetRolloutColName(), key)
}
