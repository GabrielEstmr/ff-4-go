package ff_gateways_cacheddb

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	"log"
)

type FeatureFlagCachedMongoGatewayImpl struct {
	gateway      ff_gateways.FeatureFlagsGateway
	cacheGateway ff_gateways.FeatureFlagsGateway
}

func NewFeatureFlagCachedMongoGatewayImpl(
	gateway ff_gateways.FeatureFlagsGateway,
	cacheGateway ff_gateways.FeatureFlagsGateway,
) *FeatureFlagCachedMongoGatewayImpl {
	return &FeatureFlagCachedMongoGatewayImpl{
		gateway:      gateway,
		cacheGateway: cacheGateway,
	}
}

func (this *FeatureFlagCachedMongoGatewayImpl) Save(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	savedFeatureFlagDoc, err := this.gateway.Save(featureFlag)
	if err != nil {
		return *new(ff_domains.FeatureFlag), err
	}

	go func() {
		_, err := this.cacheGateway.Save(savedFeatureFlagDoc)
		if err != nil {
			log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
		}
	}()
	return savedFeatureFlagDoc, nil
}

func (this *FeatureFlagCachedMongoGatewayImpl) Update(
	featureFlag ff_domains.FeatureFlag,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	_, err := this.cacheGateway.Update(featureFlag)
	if err != nil {
		log.Println("ff-4-go error: error to update in cache gateway. err:", err.Error())
	}

	updatedFeatureFlagDoc, err := this.gateway.Update(featureFlag)
	if err != nil {
		return *new(ff_domains.FeatureFlag), err
	}

	return updatedFeatureFlagDoc, nil
}

func (this *FeatureFlagCachedMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errC := this.cacheGateway.Delete(key)
	if errC != nil {
		return errC
	}

	errDb := this.gateway.Delete(key)
	if errDb != nil {
		return errDb
	}

	return nil
}

func (this *FeatureFlagCachedMongoGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException) {

	featureFlagDoc, errC := this.cacheGateway.FindById(key)
	if errC == nil && !featureFlagDoc.IsEmpty() {
		return featureFlagDoc, nil
	}

	featureFlagDB, errDB := this.gateway.FindById(key)
	go func() {
		if errDB == nil && !featureFlagDB.IsEmpty() {
			_, err := this.cacheGateway.Save(featureFlagDB)
			if err != nil {
				log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
			}
		}
	}()
	return featureFlagDB, errDB
}
