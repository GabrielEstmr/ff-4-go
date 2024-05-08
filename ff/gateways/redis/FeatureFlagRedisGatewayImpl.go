package ff_mongo_redis

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_redis_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis/documents"
	ff_gateways_redis_repo "github.com/GabrielEstmr/ff-4-go/ff/gateways/redis/repo"
)

type FeatureFlagRedisGatewayImpl struct {
	repo *ff_gateways_redis_repo.FeatureFlagRedisRepo
}

func NewFeatureFlagRedisGatewayImpl(ffConfigData ff.FfClientArgs) *FeatureFlagRedisGatewayImpl {
	return &FeatureFlagRedisGatewayImpl{
		repo: ff_gateways_redis_repo.NewFeatureFlagRedisRepo(ffConfigData),
	}
}

func (this *FeatureFlagRedisGatewayImpl) Save(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	savedFeatureDoc, err := this.repo.Save(
		ff_gateways_redis_documents.NewFeatureFlagDocument(featureFlag))
	if err != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeatureFlagRedisGatewayImpl) Update(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	updatedFeatureDoc, err := this.repo.Save(
		ff_gateways_redis_documents.NewFeatureFlagDocument(featureFlag))
	if err != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedFeatureDoc.ToDomain(), nil
}

func (this *FeatureFlagRedisGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	err := this.repo.Delete(key)
	if err != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return nil
}

func (this *FeatureFlagRedisGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	featureDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if featureDoc.IsEmpty() {
		return *new(ff_domains.FeatureFlag), nil
	}

	return featureDoc.ToDomain(), nil
}
