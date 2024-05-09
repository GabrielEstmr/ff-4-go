/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package ff_gateways_mongo_repo

import (
	"context"
	"errors"
	ff_configs_resources "github.com/GabrielEstmr/ff-4-go/ff/configs/resources"
	ff_gateways_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/documents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FeaturePropertyMongoRepo struct {
	_FIELD_KEY                string
	_FIELD_GROUP              string
	_FIELD_DESCRIPTION        string
	_FIELD_ENABLED            string
	_FIELD_VALUES             string
	_FIELD_LAST_MODIFIED_DATE string
	ffConfigData              ff_configs_resources.FfClientArgs
}

func NewFeaturePropertyMongoRepo(ffConfigData ff_configs_resources.FfClientArgs) *FeaturePropertyMongoRepo {
	return &FeaturePropertyMongoRepo{
		_FIELD_KEY:                "_id",
		_FIELD_GROUP:              "group",
		_FIELD_DESCRIPTION:        "description",
		_FIELD_ENABLED:            "enabled",
		_FIELD_VALUES:             "values",
		_FIELD_LAST_MODIFIED_DATE: "last_modified_date",
		ffConfigData:              ffConfigData,
	}
}

func (this *FeaturePropertyMongoRepo) Save(
	property ff_gateways_mongo_documents.FeaturePropertyDocument,
) (ff_gateways_mongo_documents.FeaturePropertyDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	property.CreatedDate = now
	property.LastModifiedDate = now

	result, err := collection.InsertOne(context.TODO(), property)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeaturePropertyDocument), err
	}

	key, _ := result.InsertedID.(string)
	property.Key = key
	return property, nil
}

func (this *FeaturePropertyMongoRepo) Update(
	property ff_gateways_mongo_documents.FeaturePropertyDocument,
) (ff_gateways_mongo_documents.FeaturePropertyDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	property.LastModifiedDate = now

	filter := bson.D{{this._FIELD_KEY, property.Key}}
	update := bson.D{
		{"$set", bson.D{{this._FIELD_GROUP, property.Group}}},
		{"$set", bson.D{{this._FIELD_DESCRIPTION, property.Description}}},
		{"$set", bson.D{{this._FIELD_ENABLED, property.Enabled}}},
		{"$set", bson.D{{this._FIELD_VALUES, property.Values}}},
		{"$set", bson.D{{this._FIELD_LAST_MODIFIED_DATE, property.LastModifiedDate}}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeaturePropertyDocument), err
	}
	return property, nil
}

func (this *FeaturePropertyMongoRepo) Delete(key string) error {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	filter := bson.D{{this._FIELD_KEY, key}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (this *FeaturePropertyMongoRepo) FindById(id string) (*ff_gateways_mongo_documents.FeaturePropertyDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	var result ff_gateways_mongo_documents.FeaturePropertyDocument
	filter := bson.D{{this._FIELD_KEY, id}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err2
	}
	return &result, nil
}
