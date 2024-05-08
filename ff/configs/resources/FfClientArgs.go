package ff_configs_resources

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// FfClientArgs is the property for the client initialization.
type FfClientArgs struct {
	db                  *mongo.Database
	dbType              DbType
	caching             bool
	cacheClient         *redis.Client
	cachingPrefix       string
	cacheType           CacheType
	featuresFlagColName string
	featuresColName     string
	rolloutColName      string
	featureFlags        ff_domains.FeatureFlags
	featureProperties   ff_domains.FeatureProperties
	rollouts            ff_domains.Rollouts
	baseRoutePath       string
}

// NewMongoFfConfigData is the constructor of FfClientArgs based on MongoDb client
// by using this method, the FfClientArgs is instantiated with default values shown below:
//   - dbType:              MONGO
//   - featuresFlagColName: "ff-feature-flags"
//   - featuresColName:     "ff-features"
//   - rolloutColName:      "ff-rollouts"
//   - baseRoutePath:       "/ff-go"
//   - caching:             false
func NewMongoFfConfigData(
	db *mongo.Database) *FfClientArgs {
	return &FfClientArgs{
		db:                  db,
		dbType:              MONGO,
		featuresFlagColName: "ff-feature-flags",
		featuresColName:     "ff-features",
		rolloutColName:      "ff-rollouts",
		baseRoutePath:       "/ff",
		caching:             false,
	}
}

// WithRedisCache add the RedisClient in FfClientArgs
func (this FfClientArgs) WithRedisCache(redisClient *redis.Client) FfClientArgs {
	this.caching = true
	this.cacheClient = redisClient
	this.cachingPrefix = "ff-4-go"
	this.cacheType = REDIS
	return this
}

// WithCustomFeatureFlagColName changes the name of the collection name for FeatureFlags persistence in main Database
// by default, this value is given by NewMongoFfConfigData and is configured as ff-feature-flags
func (this FfClientArgs) WithCustomFeatureFlagColName(featuresFlagColName string) FfClientArgs {
	this.featuresFlagColName = featuresFlagColName
	return this
}

// WithCustomFeatureColName changes the name of the collection name for FeatureProperties persistence in main Database
// by default, this value is given by NewMongoFfConfigData and is configured as ff-features
func (this FfClientArgs) WithCustomFeatureColName(featuresColName string) FfClientArgs {
	this.featuresColName = featuresColName
	return this
}

// WithCustomRolloutColName changes the name of the collection name for FeatureProperties persistence in main Database
// This value is used to build the key for the cache implementation as well
// by default, this value is given by NewMongoFfConfigData and is configured as ff-rollouts
func (this FfClientArgs) WithCustomRolloutColName(rolloutColName string) FfClientArgs {
	this.rolloutColName = rolloutColName
	return this
}

// WithCustomCachePrefix changes the caching prefix to avoid multiple applications key conflicts
// by default, this value is given by WithRedisCache and is configured as ff-4-go
func (this FfClientArgs) WithCustomCachePrefix(cachingPrefix string) FfClientArgs {
	this.cachingPrefix = cachingPrefix
	return this
}

// WithFeatureFlagsInitialValues adds new values to feature-flags when the application starts
// If this property is configured as nil, the application will not create the FeatureFlags
// by default, this value is given by NewMongoFfConfigData and is configured as nil
func (this FfClientArgs) WithFeatureFlagsInitialValues(featureFlags ff_domains.FeatureFlags,
) FfClientArgs {
	this.featureFlags = featureFlags
	return this
}

// WithFeaturePropertiesInitialValues adds new values to feature-properties when the application starts
// If this property is configured as nil, the application will not create the FeatureFlags
// by default, this value is given by NewMongoFfConfigData and is configured as nil
func (this FfClientArgs) WithFeaturePropertiesInitialValues(featureProperties ff_domains.FeatureProperties,
) FfClientArgs {
	this.featureProperties = featureProperties
	return this
}

// WithRolloutsInitialValues adds new values to rollouts when the application starts
// If this property is configured as nil, the application will not create the FeatureFlags
// by default, this value is given by NewMongoFfConfigData and is configured as nil
func (this FfClientArgs) WithRolloutsInitialValues(rollouts ff_domains.Rollouts,
) FfClientArgs {
	this.rollouts = rollouts
	return this
}

// WithCustomBaseRoutePath changes the base route for the endpoints given by the library
// by default, this value is given by NewMongoFfConfigData and is configured as /ff-go
func (this FfClientArgs) WithCustomBaseRoutePath(baseRoutePath string) FfClientArgs {
	this.baseRoutePath = baseRoutePath
	return this
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

func (this FfClientArgs) GetFeaturesFlagColName() string {
	return this.featuresFlagColName
}

func (this FfClientArgs) GetFeaturesColName() string {
	return this.featuresColName
}

func (this FfClientArgs) GetRolloutColName() string {
	return this.rolloutColName
}

func (this FfClientArgs) GetFeatureFlags() ff_domains.FeatureFlags {
	return this.featureFlags
}

func (this FfClientArgs) GetFeatureProperties() ff_domains.FeatureProperties {
	return this.featureProperties
}

func (this FfClientArgs) GetRollouts() ff_domains.Rollouts {
	return this.rollouts
}

func (this FfClientArgs) GetBaseRoutePath() string {
	return this.baseRoutePath
}

func (this FfClientArgs) IsMongoType() bool {
	return this.dbType == MONGO
}

func (this FfClientArgs) IsRedisCacheType() bool {
	return this.cacheType == REDIS
}

func (this FfClientArgs) HasInitialFeatureFlagValues() bool {
	return this.featureFlags != nil
}

func (this FfClientArgs) HasInitialFeaturePropertiesValues() bool {
	return this.featureProperties != nil
}

func (this FfClientArgs) HasInitialRolloutsValues() bool {
	return this.rollouts != nil
}
