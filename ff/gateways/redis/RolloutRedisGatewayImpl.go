package ff_mongo_redis

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_redis_documents "baseapplicationgo/main/configs/ff/lib/gateways/redis/documents"
	ff_gateways_redis_repo "baseapplicationgo/main/configs/ff/lib/gateways/redis/repo"
)

type RolloutRedisGatewayImpl struct {
	repo *ff_gateways_redis_repo.RolloutRedisRepo
}

func NewRolloutRedisGatewayImpl(ffConfigData ff.FfClientArgs) *RolloutRedisGatewayImpl {
	return &RolloutRedisGatewayImpl{
		repo: ff_gateways_redis_repo.NewRolloutRedisRepo(ffConfigData),
	}
}

func (this *RolloutRedisGatewayImpl) Save(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	savedRolloutDoc, err := this.repo.Save(
		*ff_gateways_redis_documents.NewRolloutDocument(rollout))
	if err != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedRolloutDoc.ToDomain(), nil
}

func (this *RolloutRedisGatewayImpl) Update(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	updatedRolloutDoc, err := this.repo.Save(
		*ff_gateways_redis_documents.NewRolloutDocument(rollout))
	if err != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedRolloutDoc.ToDomain(), nil
}

func (this *RolloutRedisGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	err := this.repo.Delete(key)
	if err != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return nil
}

func (this *RolloutRedisGatewayImpl) FindById(
	key string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	featureDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if featureDoc.IsEmpty() {
		return *new(ff_domains.Rollout), nil
	}

	return featureDoc.ToDomain(), nil
}
