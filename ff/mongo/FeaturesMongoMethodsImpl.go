package ff_mongo

import (
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_mongo_exceptions "github.com/GabrielEstmr/ff-4-go/ff/mongo/exceptions"
	ff_mongo_repo "github.com/GabrielEstmr/ff-4-go/ff/mongo/repo"
	ff_resources "github.com/GabrielEstmr/ff-4-go/ff/resources"
)

type FeaturesMongoMethodsImpl struct {
	repo *ff_mongo_repo.FeaturesMongoRepo
}

func NewFeaturesMongoMethodsImpl(ffConfigData *ff.FfConfigData) *FeaturesMongoMethodsImpl {
	return &FeaturesMongoMethodsImpl{repo: ff_mongo_repo.NewFeaturesMongoRepo(ffConfigData)}
}

func (this *FeaturesMongoMethodsImpl) getFeature(key string) (
	ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}

	if byId.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	return byId.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) IsEnabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsEnabled(), nil
}

func (this *FeaturesMongoMethodsImpl) IsDisabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}

func (this *FeaturesMongoMethodsImpl) Enable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	if featureDoc.IsDisabled() {
		featureDoc.DefaultValue = true
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(ff_resources.FeaturesData),
				ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}

func (this *FeaturesMongoMethodsImpl) Disable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	if featureDoc.IsEnabled() {
		featureDoc.DefaultValue = false
		savedFeatureDoc, err := this.repo.Update(*featureDoc)
		if err != nil {
			return *new(ff_resources.FeaturesData),
				ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
		}
		return savedFeatureDoc.ToDomain(), nil
	}
	return featureDoc.ToDomain(), nil
}
