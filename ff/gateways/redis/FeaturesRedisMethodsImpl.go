package ff_mongo_redis

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_redis_documents "baseapplicationgo/main/configs/ff/lib/gateways/redis/documents"
	ff_gateways_redis_repo "baseapplicationgo/main/configs/ff/lib/gateways/redis/repo"
)

type FeaturesRedisGatewayImpl struct {
	repo *ff_gateways_redis_repo.FeaturesRedisRepo
}

func NewFeaturesRedisGatewayImpl(ffConfigData ff.FfClientArgs) *FeaturesRedisGatewayImpl {
	return &FeaturesRedisGatewayImpl{
		repo: ff_gateways_redis_repo.NewFeaturesRedisRepo(ffConfigData),
	}
}

func (this *FeaturesRedisGatewayImpl) Save(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	savedFeatureDoc, err := this.repo.Save(
		ff_gateways_redis_documents.NewFeaturesDataRedisDocument(feature))
	if err != nil {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesRedisGatewayImpl) Update(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	updatedFeatureDoc, err := this.repo.Save(
		ff_gateways_redis_documents.NewFeaturesDataRedisDocument(feature))
	if err != nil {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesRedisGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	err := this.repo.Delete(key)
	if err != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return nil
}

func (this *FeaturesRedisGatewayImpl) FindById(
	key string,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	featureDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if featureDoc.IsEmpty() {
		return *new(ff_domains.Feature), nil
	}

	return featureDoc.ToDomain(), nil
}
