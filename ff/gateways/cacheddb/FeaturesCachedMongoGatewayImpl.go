package ff_gateways_cacheddb

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways "baseapplicationgo/main/configs/ff/lib/gateways"
	"log"
)

type FeaturesCachedMongoGatewayImpl struct {
	gateway      ff_gateways.FeaturesGateway
	cacheGateway ff_gateways.FeaturesGateway
}

func NewFeaturesCachedMongoGatewayImpl(
	gateway ff_gateways.FeaturesGateway,
	cacheGateway ff_gateways.FeaturesGateway,
) *FeaturesCachedMongoGatewayImpl {
	return &FeaturesCachedMongoGatewayImpl{
		gateway:      gateway,
		cacheGateway: cacheGateway,
	}
}

func (this *FeaturesCachedMongoGatewayImpl) Save(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	savedFeatureDoc, err := this.gateway.Save(feature)
	if err != nil {
		return *new(ff_domains.Feature), err
	}

	go func() {
		_, err := this.cacheGateway.Save(feature)
		if err != nil {
			log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
		}
	}()
	return savedFeatureDoc, nil
}

func (this *FeaturesCachedMongoGatewayImpl) Update(
	feature ff_domains.Feature,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	updatedFeatureDoc, err := this.gateway.Update(feature)
	if err != nil {
		return *new(ff_domains.Feature), err
	}

	go func() {
		_, err := this.cacheGateway.Update(feature)
		if err != nil {
			log.Println("ff-4-go error: error to update in cache gateway. err:", err.Error())
		}
	}()
	return updatedFeatureDoc, nil
}

func (this *FeaturesCachedMongoGatewayImpl) Delete(
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

func (this *FeaturesCachedMongoGatewayImpl) FindById(
	key string,
) (ff_domains.Feature, ff_domains_exceptions.LibException) {

	featureDoc, errC := this.cacheGateway.FindById(key)
	if errC == nil && !featureDoc.IsEmpty() {
		return featureDoc, nil
	}

	return this.gateway.FindById(key)
}
