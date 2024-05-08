package ff_gateways_beans

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_gateways_cacheddb "baseapplicationgo/main/configs/ff/lib/gateways/cacheddb"
	ff_gateways_mongo "baseapplicationgo/main/configs/ff/lib/gateways/mongo"
	ff_mongo_redis "baseapplicationgo/main/configs/ff/lib/gateways/redis"
	"errors"
)

type FeatureFlagGatewayFactory struct {
	ffConfigData ff.FfClientArgs
}

func NewFeatureFlagGatewayFactory(ffConfigData ff.FfClientArgs) *FeatureFlagGatewayFactory {
	return &FeatureFlagGatewayFactory{ffConfigData}
}

func (this *FeatureFlagGatewayFactory) Get() (ff_gateways.FeatureFlagsGateway, error) {
	if this.ffConfigData.IsMongoType() &&
		!this.ffConfigData.HasCaching() {
		return ff_gateways_mongo.NewFeatureFlagMongoGatewayImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.IsMongoType() &&
		this.ffConfigData.HasCaching() &&
		this.ffConfigData.IsRedisCacheType() {
		return ff_gateways_cacheddb.NewFeatureFlagCachedMongoGatewayImpl(
			ff_gateways_mongo.NewFeatureFlagMongoGatewayImpl(this.ffConfigData),
			ff_mongo_redis.NewFeatureFlagRedisGatewayImpl(this.ffConfigData),
		), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesGateway")
}
