/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_mongo

import (
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/documents"
	ff_gateways_mongo_repo "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/repo"
)

type FeaturePropertyMongoGatewayImpl struct {
	repo *ff_gateways_mongo_repo.FeaturePropertyMongoRepo
}

func NewFeaturePropertyMongoGatewayImpl(ffConfigData ff_configs_resources.FfClientArgs) *FeaturePropertyMongoGatewayImpl {
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
