package ff_gateways_beans

import (
	"errors"
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_gateways_cacheddb "github.com/GabrielEstmr/ff-4-go/ff/gateways/cacheddb"
	ff_gateways_mongo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo"
	ff_mongo_redis "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis"
)

type FeatureFlagGatewayFactory struct {
	ffConfigData ff_configs_resources.FfClientArgs
}

func NewFeatureFlagGatewayFactory(ffConfigData ff_configs_resources.FfClientArgs) *FeatureFlagGatewayFactory {
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
