package ff_gateways_beans

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_gateways_cacheddb "baseapplicationgo/main/configs/ff/lib/gateways/cacheddb"
	ff_gateways_mongo "baseapplicationgo/main/configs/ff/lib/gateways/mongo"
	ff_mongo_redis "baseapplicationgo/main/configs/ff/lib/gateways/redis"
	"errors"
)

type FeaturePropertyGatewayFactory struct {
	ffConfigData ff.FfClientArgs
}

func NewFeaturePropertyGatewayFactory(ffConfigData ff.FfClientArgs) *FeaturePropertyGatewayFactory {
	return &FeaturePropertyGatewayFactory{ffConfigData}
}

func (this *FeaturePropertyGatewayFactory) Get() (ff_gateways.FeaturePropertyGateway, error) {
	if this.ffConfigData.IsMongoType() &&
		!this.ffConfigData.HasCaching() {
		return ff_gateways_mongo.NewFeaturePropertyMongoGatewayImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.IsMongoType() &&
		this.ffConfigData.HasCaching() &&
		this.ffConfigData.IsRedisCacheType() {
		return ff_gateways_cacheddb.NewFeaturePropertyCachedMongoGatewayImpl(
			ff_gateways_mongo.NewFeaturePropertyMongoGatewayImpl(this.ffConfigData),
			ff_mongo_redis.NewFeaturePropertyRedisGatewayImpl(this.ffConfigData),
		), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesPropertyGateway")
}
