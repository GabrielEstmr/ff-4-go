package ff_gateways_mongo

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_mongo_documents "baseapplicationgo/main/configs/ff/lib/gateways/mongo/documents"
	ff_gateways_mongo_repo "baseapplicationgo/main/configs/ff/lib/gateways/mongo/repo"
)

type FeaturePropertyMongoGatewayImpl struct {
	repo *ff_gateways_mongo_repo.FeaturePropertyMongoRepo
}

func NewFeaturePropertyMongoGatewayImpl(ffConfigData ff.FfClientArgs) *FeaturePropertyMongoGatewayImpl {
	return &FeaturePropertyMongoGatewayImpl{
		repo: ff_gateways_mongo_repo.NewFeaturePropertyMongoRepo(ffConfigData),
	}
}

func (this *FeaturePropertyMongoGatewayImpl) Save(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	savedPropertyDoc, err := this.repo.Save(
		*ff_gateways_mongo_documents.NewFeaturePropertyDocument(property))
	if err != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedPropertyDoc.ToDomain(), nil
}

func (this *FeaturePropertyMongoGatewayImpl) Update(
	property ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	updatedPropertyDoc, err := this.repo.Update(
		*ff_gateways_mongo_documents.NewFeaturePropertyDocument(property))
	if err != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return updatedPropertyDoc.ToDomain(), nil
}

func (this *FeaturePropertyMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errDb := this.repo.Delete(key)
	if errDb != nil {
		return ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}
	return nil
}

func (this *FeaturePropertyMongoGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	propertyDoc, errDb := this.repo.FindById(key)
	if errDb != nil {
		return *new(ff_domains.FeatureProperty),
			ff_domains_exceptions.NewInternalServerErrorExceptionSglMsg(errDb.Error())
	}

	if propertyDoc.IsEmpty() {
		return *new(ff_domains.FeatureProperty), nil
	}

	return propertyDoc.ToDomain(), nil
}
