package ff_gateways_mongo

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_mongo_documents "baseapplicationgo/main/configs/ff/lib/gateways/mongo/documents"
	ff_gateways_mongo_repo "baseapplicationgo/main/configs/ff/lib/gateways/mongo/repo"
)

type FeaturesMongoGatewayImpl struct {
	repo *ff_gateways_mongo_repo.FeaturesMongoRepo
}

func NewFeaturesMongoMethodsImpl(ffConfigData ff.FfClientArgs) *FeaturesMongoGatewayImpl {
	return &FeaturesMongoGatewayImpl{
		repo: ff_gateways_mongo_repo.NewFeaturesMongoRepo(ffConfigData),
	}
}

func (this *FeaturesMongoGatewayImpl) Save(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	savedFeatureDoc, err := this.repo.Save(
		ff_gateways_mongo_documents.NewFeaturesDataDocument(feature))
	if err != nil {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesMongoGatewayImpl) Update(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	updatedFeatureDoc, err := this.repo.Update(
		ff_gateways_mongo_documents.NewFeaturesDataDocument(feature))
	if err != nil {
		return *new(ff_domains.Feature),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errDb := this.repo.Delete(key)
	if errDb != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}
	return nil
}

func (this *FeaturesMongoGatewayImpl) FindById(
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
