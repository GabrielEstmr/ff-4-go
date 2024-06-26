/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_cacheddb

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	"log"
)

type FeaturePropertyCachedMongoGatewayImpl struct {
	gateway      ff_gateways.FeaturePropertyGateway
	cacheGateway ff_gateways.FeaturePropertyGateway
}

func NewFeaturePropertyCachedMongoGatewayImpl(
	gateway ff_gateways.FeaturePropertyGateway,
	cacheGateway ff_gateways.FeaturePropertyGateway,
) *FeaturePropertyCachedMongoGatewayImpl {
	return &FeaturePropertyCachedMongoGatewayImpl{
		gateway:      gateway,
		cacheGateway: cacheGateway,
	}
}

func (this *FeaturePropertyCachedMongoGatewayImpl) Save(
	feature ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	savedFeature, err := this.gateway.Save(feature)
	if err != nil {
		return *new(ff_domains.FeatureProperty), err
	}

	go func() {
		_, err := this.cacheGateway.Save(savedFeature)
		if err != nil {
			log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
		}
	}()
	return savedFeature, nil
}

func (this *FeaturePropertyCachedMongoGatewayImpl) Update(
	feature ff_domains.FeatureProperty,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	_, err := this.cacheGateway.Update(feature)
	if err != nil {
		log.Println("ff-4-go error: error to update in cache gateway. err:", err.Error())
	}

	updatedFeatureDoc, err := this.gateway.Update(feature)
	if err != nil {
		return *new(ff_domains.FeatureProperty), err
	}

	return updatedFeatureDoc, nil
}

func (this *FeaturePropertyCachedMongoGatewayImpl) Delete(
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

func (this *FeaturePropertyCachedMongoGatewayImpl) FindById(
	key string,
) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException) {

	feature, errC := this.cacheGateway.FindById(key)
	if errC == nil && !feature.IsEmpty() {
		return feature, nil
	}

	featureDB, errDB := this.gateway.FindById(key)
	go func() {
		if errDB == nil && !featureDB.IsEmpty() {
			_, err := this.cacheGateway.Save(featureDB)
			if err != nil {
				log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
			}
		}
	}()
	return featureDB, errDB
}
