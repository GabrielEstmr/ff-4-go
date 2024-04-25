package ff_mongo

import (
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_mongo_exceptions "github.com/GabrielEstmr/ff-4-go/ff/mongo/exceptions"
	ff_mongo_repo "github.com/GabrielEstmr/ff-4-go/ff/mongo/repo"
	ff_redis_documents "github.com/GabrielEstmr/ff-4-go/ff/redis/documents"
	ff_redis_repo "github.com/GabrielEstmr/ff-4-go/ff/redis/repo"
	ff_resources "github.com/GabrielEstmr/ff-4-go/ff/resources"
)

type FeaturesCachedMongoMethodsImpl struct {
	repo      *ff_mongo_repo.FeaturesMongoRepo
	repoCache *ff_redis_repo.FeaturesRedisRepo
}

func NewFeaturesCachedMongoMethodsImpl(ffConfigData *ff.FfConfigData) *FeaturesCachedMongoMethodsImpl {
	return &FeaturesCachedMongoMethodsImpl{
		repo:      ff_mongo_repo.NewFeaturesMongoRepo(ffConfigData),
		repoCache: ff_redis_repo.NewFeaturesRedisRepo(ffConfigData),
	}
}

func (this *FeaturesCachedMongoMethodsImpl) getFeature(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {

	byIdCached, err := this.repoCache.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	if !byIdCached.IsEmpty() {
		return byIdCached.ToDomain(), nil
	}

	byId, err := this.repo.FindById(key)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}

	if byId.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	go this.repoCache.Save(ff_redis_documents.NewFeaturesDataRedisDocument(byId.ToDomain()))
	return byId.ToDomain(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) IsEnabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsEnabled(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) IsDisabled(key string,
) (bool, ff_mongo_exceptions.LibException) {
	feature, err := this.getFeature(key)
	if err != nil {
		return false, err
	}
	return feature.IsDisabled(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) Enable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, errFind := this.repo.FindById(key)
	if errFind != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(errFind.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	featureDoc.DefaultValue = true
	_, errCache := this.repoCache.Save(ff_redis_documents.NewFeaturesDataRedisDocument(featureDoc.ToDomain()))
	if errCache != nil {
		return *new(ff_resources.FeaturesData), ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(errCache.Error())
	}

	savedFeatureDoc, err := this.repo.Update(*featureDoc)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}

func (this *FeaturesCachedMongoMethodsImpl) Disable(key string,
) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException) {
	featureDoc, errFind := this.repo.FindById(key)
	if errFind != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(errFind.Error())
	}
	if featureDoc.IsEmpty() {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibResourceNotFoundExceptionSglMsg("feature not found")
	}

	featureDoc.DefaultValue = false
	_, errCache := this.repoCache.Save(ff_redis_documents.NewFeaturesDataRedisDocument(featureDoc.ToDomain()))
	if errCache != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(errCache.Error())
	}

	savedFeatureDoc, err := this.repo.Update(*featureDoc)
	if err != nil {
		return *new(ff_resources.FeaturesData),
			ff_mongo_exceptions.NewLibInternalServerErrorExceptionSglMsg(err.Error())
	}
	return savedFeatureDoc.ToDomain(), nil
}
