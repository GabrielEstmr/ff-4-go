package ff_gateways_mongo

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/documents"
	ff_gateways_mongo_repo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/repo"
)

type RolloutMongoGatewayImpl struct {
	repo *ff_gateways_mongo_repo.RolloutMongoRepo
}

func NewRolloutMongoGatewayImpl(ffConfigData ff.FfClientArgs) *RolloutMongoGatewayImpl {
	return &RolloutMongoGatewayImpl{
		repo: ff_gateways_mongo_repo.NewRolloutMongoRepo(ffConfigData),
	}
}

func (this *RolloutMongoGatewayImpl) Save(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	savedRolloutDoc, err := this.repo.Save(
		*ff_gateways_mongo_documents.NewRolloutDocument(rollout))
	if err != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedRolloutDoc.ToDomain(), nil
}

func (this *RolloutMongoGatewayImpl) Update(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	updatedRolloutDoc, err := this.repo.Update(
		*ff_gateways_mongo_documents.NewRolloutDocument(rollout))
	if err != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedRolloutDoc.ToDomain(), nil
}

func (this *RolloutMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errDb := this.repo.Delete(key)
	if errDb != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}
	return nil
}

func (this *RolloutMongoGatewayImpl) FindById(
	key string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	rolloutDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.Rollout),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if rolloutDoc.IsEmpty() {
		return *new(ff_domains.Rollout), nil
	}

	return rolloutDoc.ToDomain(), nil
}
