package ff_factories

import (
	"errors"
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_mongo "github.com/GabrielEstmr/ff-4-go/ff/mongo"
)

type FeaturesMethodsFactory struct {
	ffConfigData *ff.FfConfigData
}

func NewFeaturesMethodsFactory(ffConfigData *ff.FfConfigData) *FeaturesMethodsFactory {
	return &FeaturesMethodsFactory{ffConfigData}
}

func (this *FeaturesMethodsFactory) Get() (ff.FeaturesMethods, error) {
	if this.ffConfigData.GetClientType() == ff.MONGO &&
		this.ffConfigData.GetHasCaching() == false {
		return ff_mongo.NewFeaturesMongoMethodsImpl(this.ffConfigData), nil
	}

	if this.ffConfigData.GetClientType() == ff.MONGO &&
		this.ffConfigData.GetHasCaching() == true &&
		this.ffConfigData.GetCacheClientType() == ff.REDIS {
		return ff_mongo.NewFeaturesCachedMongoMethodsImpl(this.ffConfigData), nil
	}

	return nil, errors.New("could not instantiate a valid FeaturesData")
}
