package ff_gateways_beans

import (
	"errors"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_gateways_cacheddb "github.com/GabrielEstmr/ff-4-go/ff/gateways/cacheddb"
	ff_gateways_mongo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo"
	ff_mongo_redis "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis"
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
