package ff_mongo_redis

import (
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_redis_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis/documents"
	ff_gateways_redis_repo "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis/repo"
)

type FeaturePropertyRedisGatewayImpl struct {
	repo *ff_gateways_redis_repo.FeaturePropertyRedisRepo
}

func NewFeaturePropertyRedisGatewayImpl(ffConfigData ff_configs_resources.FfClientArgs) *FeaturePropertyRedisGatewayImpl {
	return &FeaturePropertyRedisGatewayImpl{
		repo: ff_gateways_redis_repo.NewFeaturePropertyRedisRepo(ffConfigData),
	}
}

func (this *FeaturePropertyRedisGatewayImpl) Save(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	savedRolloutDoc, err := this.repo.Save(
		*ff_gateways_redis_documents.NewFeaturePropertyDocument(property))
	if err != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedRolloutDoc.ToDomain(), nil
}

func (this *FeaturePropertyRedisGatewayImpl) Update(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	updatedRolloutDoc, err := this.repo.Save(
		*ff_gateways_redis_documents.NewFeaturePropertyDocument(property))
	if err != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedRolloutDoc.ToDomain(), nil
}

func (this *FeaturePropertyRedisGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	err := this.repo.Delete(key)
	if err != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return nil
}

func (this *FeaturePropertyRedisGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	featureDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if featureDoc.IsEmpty() {
		return *new(ff_domains.FeatureProperty), nil
	}

	return featureDoc.ToDomain(), nil
}
