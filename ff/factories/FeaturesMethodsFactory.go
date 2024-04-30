package ff_factories

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_gateways_cacheddb "baseapplicationgo/main/configs/ff/lib/gateways/cacheddb"
	ff_gateways_mongo "baseapplicationgo/main/configs/ff/lib/gateways/mongo"
	ff_mongo_redis "baseapplicationgo/main/configs/ff/lib/gateways/redis"
	"errors"
)

type FeaturesGatewayFactory struct {
	ffConfigData ff.FfClientArgs
}

func NewFeaturesGatewayFactory(ffConfigData ff.FfClientArgs) *FeaturesGatewayFactory {
	return &FeaturesGatewayFactory{ffConfigData}
}

func (this *FeaturesGatewayFactory) Get() (ff_gateways.FeaturesGateway, error) {
	if this.ffConfigData.IsMongoType() &&
		!this.ffConfigData.HasCaching() {
		return ff_gateways_mongo.NewFeaturesMongoMethodsImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.IsMongoType() &&
		this.ffConfigData.HasCaching() &&
		this.ffConfigData.IsRedisCacheType() {
		return ff_gateways_cacheddb.NewFeaturesCachedMongoGatewayImpl(
			ff_gateways_mongo.NewFeaturesMongoMethodsImpl(this.ffConfigData),
			ff_mongo_redis.NewFeaturesRedisGatewayImpl(this.ffConfigData),
		), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesData")
}
