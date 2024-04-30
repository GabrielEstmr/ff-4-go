package ff_configs_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type FfClientArgs struct {
	db                    *mongo.Database
	dbType                DbType
	caching               bool
	cacheClient           *redis.Client
	cachingPrefix         string
	cacheType             CacheType
	featuresColName       string
	featuresInitialValues ff_domains.Features
	baseRoutePath         string
}

func NewMongoFfConfigData(
	db *mongo.Database) *FfClientArgs {
	return &FfClientArgs{
		db:              db,
		dbType:          MONGO,
		featuresColName: "ff-features",
		baseRoutePath:   "/ff-go",
		caching:         false,
	}
}

func (this FfClientArgs) GetDb() *mongo.Database {
	return this.db
}

func (this FfClientArgs) GetDbType() DbType {
	return this.dbType
}

func (this FfClientArgs) GetCacheClient() *redis.Client {
	return this.cacheClient
}

func (this FfClientArgs) HasCaching() bool {
	return this.caching
}

func (this FfClientArgs) GetCachingPrefix() string {
	return this.cachingPrefix
}

func (this FfClientArgs) GetCacheClientType() CacheType {
	return this.cacheType
}

func (this FfClientArgs) GetFeaturesColName() string {
	return this.featuresColName
}

func (this FfClientArgs) GetFeaturesInitialValues() ff_domains.Features {
	return this.featuresInitialValues
}

func (this FfClientArgs) GetBaseRoutePath() string {
	return this.baseRoutePath
}

func (this FfClientArgs) WithRedisCache(redisClient *redis.Client) FfClientArgs {
	this.caching = true
	this.cacheClient = redisClient
	this.cachingPrefix = "ff-4-go"
	this.cacheType = REDIS
	return this
}

func (this FfClientArgs) WithPersonalizedColName(featuresColName string) FfClientArgs {
	this.featuresColName = featuresColName
	return this
}

func (this FfClientArgs) WithPersonalizedCachePrefix(cachingPrefix string) FfClientArgs {
	this.cachingPrefix = cachingPrefix
	return this
}

func (this FfClientArgs) WithFeatureInitialValues(featuresInitialValues ff_domains.Features,
) FfClientArgs {
	this.featuresInitialValues = featuresInitialValues
	return this
}

func (this FfClientArgs) WithPersonalizedBaseRoutePath(baseRoutePath string) FfClientArgs {
	this.baseRoutePath = baseRoutePath
	return this
}

func (this FfClientArgs) IsMongoType() bool {
	return this.dbType == MONGO
}

func (this FfClientArgs) IsRedisCacheType() bool {
	return this.cacheType == REDIS
}

func (this FfClientArgs) HasInitialValues() bool {
	return this.featuresInitialValues != nil
}
