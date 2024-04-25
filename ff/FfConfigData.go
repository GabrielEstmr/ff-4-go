package ff

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type FfConfigData struct {
	db              *mongo.Database
	clientType      string
	hasCaching      bool
	cacheClient     *redis.Client
	cachingPrefix   string
	cacheClientType string
	featuresDbName  string
}

func NewFfConfigData(
	db *mongo.Database,
	clientType string,
	hasCaching bool,
	cacheClient *redis.Client,
	cachingPrefix string,
	cacheClientType string,
	featuresDbName string) *FfConfigData {
	return &FfConfigData{
		db:              db,
		clientType:      clientType,
		hasCaching:      hasCaching,
		cacheClient:     cacheClient,
		cachingPrefix:   cachingPrefix,
		cacheClientType: cacheClientType,
		featuresDbName:  featuresDbName}
}

func (this FfConfigData) GetDb() *mongo.Database {
	return this.db
}

func (this FfConfigData) GetClientType() string {
	return this.clientType
}

func (this FfConfigData) GetCacheClient() *redis.Client {
	return this.cacheClient
}

func (this FfConfigData) GetHasCaching() bool {
	return this.hasCaching
}

func (this FfConfigData) GetCachingPrefix() string {
	return this.cachingPrefix
}

func (this FfConfigData) GetCacheClientType() string {
	return this.cacheClientType
}

func (this FfConfigData) GetFeaturesDbName() string {
	return this.featuresDbName
}
