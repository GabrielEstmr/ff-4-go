package ff_gateways_beans

import (
	"errors"
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	ff_gateways_cacheddb "github.com/GabrielEstmr/ff-4-go/ff/gateways/cacheddb"
	ff_gateways_mongo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo"
	ff_mongo_redis "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis"
)

type RolloutGatewayFactory struct {
	ffConfigData ff_configs_resources.FfClientArgs
}

func NewRolloutGatewayFactory(ffConfigData ff_configs_resources.FfClientArgs) *RolloutGatewayFactory {
	return &RolloutGatewayFactory{ffConfigData}
}

func (this *RolloutGatewayFactory) Get() (ff_gateways.RolloutGateway, error) {
	if this.ffConfigData.IsMongoType() &&
		!this.ffConfigData.HasCaching() {
		return ff_gateways_mongo.NewRolloutMongoGatewayImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.IsMongoType() &&
		this.ffConfigData.HasCaching() &&
		this.ffConfigData.IsRedisCacheType() {
		return ff_gateways_cacheddb.NewRolloutCachedMongoGatewayImpl(
			ff_gateways_mongo.NewRolloutMongoGatewayImpl(this.ffConfigData),
			ff_mongo_redis.NewRolloutRedisGatewayImpl(this.ffConfigData),
		), nil
	}

	return nil, errors.New("could not instantiate a valid RolloutGateway")
}
