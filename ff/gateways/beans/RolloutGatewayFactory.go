package ff_gateways_beans

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	ff_gateways_cacheddb "baseapplicationgo/main/configs/ff/lib/gateways/cacheddb"
	ff_gateways_mongo "baseapplicationgo/main/configs/ff/lib/gateways/mongo"
	ff_mongo_redis "baseapplicationgo/main/configs/ff/lib/gateways/redis"
	"errors"
)

type RolloutGatewayFactory struct {
	ffConfigData ff.FfClientArgs
}

func NewRolloutGatewayFactory(ffConfigData ff.FfClientArgs) *RolloutGatewayFactory {
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
