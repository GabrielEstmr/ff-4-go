package ff

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var once sync.Once
var ffLibConfigsBean *FfConfigData

func NewFfConfigDataBean(
	db *mongo.Database,
	clientType string,
	hasCaching bool,
	cacheClient *redis.Client,
	cachingPrefix string,
	cacheClientType string,
	featuresDbName string,
) *FfConfigData {
	once.Do(func() {
		if ffLibConfigsBean == nil {
			ffLibConfigsBean = getFfConfigData(
				db,
				clientType,
				hasCaching,
				cacheClient,
				cachingPrefix,
				cacheClientType,
				featuresDbName)
		}
	})
	return ffLibConfigsBean
}

func getFfConfigData(
	db *mongo.Database,
	clientType string,
	hasCaching bool,
	cacheClient *redis.Client,
	cachingPrefix string,
	cacheClientType string,
	featuresDbName string) *FfConfigData {
	return NewFfConfigData(db,
		clientType,
		hasCaching,
		cacheClient,
		cachingPrefix,
		cacheClientType,
		featuresDbName)
}
