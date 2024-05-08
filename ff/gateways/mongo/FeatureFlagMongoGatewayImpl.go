package ff_gateways_mongo

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/documents"
	ff_gateways_mongo_repo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/repo"
)

type FeatureFlagMongoGatewayImpl struct {
	repo *ff_gateways_mongo_repo.FeatureFlagMongoRepo
}

func NewFeatureFlagMongoGatewayImpl(ffConfigData ff.FfClientArgs) *FeatureFlagMongoGatewayImpl {
	return &FeatureFlagMongoGatewayImpl{
		repo: ff_gateways_mongo_repo.NewFeatureFlagMongoRepo(ffConfigData),
	}
}

func (this *FeatureFlagMongoGatewayImpl) Save(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	savedFeatureDoc, err := this.repo.Save(
		ff_gateways_mongo_documents.NewFeatureFlagDocument(featureFlag))
	if err != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeatureFlagMongoGatewayImpl) Update(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	updatedFeatureFlagDoc, err := this.repo.Update(
		ff_gateways_mongo_documents.NewFeatureFlagDocument(featureFlag))
	if err != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedFeatureFlagDoc.ToDomain(), nil
}

func (this *FeatureFlagMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errDb := this.repo.Delete(key)
	if errDb != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}
	return nil
}

func (this *FeatureFlagMongoGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	featureFlagDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.FeatureFlag),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if featureFlagDoc.IsEmpty() {
		return *new(ff_domains.FeatureFlag), nil
	}

	return featureFlagDoc.ToDomain(), nil
}
