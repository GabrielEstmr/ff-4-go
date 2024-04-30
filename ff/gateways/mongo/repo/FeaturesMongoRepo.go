package ff_gateways_mongo_repo

import (
	ff "baseapplicationgo/main/configs/ff/lib/configs/resources"
	"baseapplicationgo/main/configs/ff/lib/gateways/mongo/documents"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const _KEY = "_id"

type FeaturesMongoRepo struct {
	ffConfigData ff.FfClientArgs
}

func NewFeaturesMongoRepo(ffConfigData ff.FfClientArgs) *FeaturesMongoRepo {
	return &FeaturesMongoRepo{ffConfigData: ffConfigData}
}

func (this *FeaturesMongoRepo) Save(
	feature ff_gateways_mongo_documents.FeaturesDataDocument,
) (ff_gateways_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())

	result, err := collection.InsertOne(context.TODO(), feature)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeaturesDataDocument), err
	}

	key, _ := result.InsertedID.(string)
	feature.Key = key
	return feature, nil
}

func (this *FeaturesMongoRepo) Update(
	feature ff_gateways_mongo_documents.FeaturesDataDocument,
) (ff_gateways_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	filter := bson.D{{_KEY, feature.Key}}
	update := bson.D{{"$set", bson.D{{"value", feature.Value}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeaturesDataDocument), err
	}
	return feature, nil
}

func (this *FeaturesMongoRepo) Delete(key string) error {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	filter := bson.D{{_KEY, key}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (this *FeaturesMongoRepo) FindById(id string) (*ff_gateways_mongo_documents.FeaturesDataDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesColName())
	var result ff_gateways_mongo_documents.FeaturesDataDocument
	filter := bson.D{{_KEY, id}}
	err2 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err2 != nil {
		if errors.Is(err2, mongo.ErrNoDocuments) {
			return &result, nil
		}
		return &result, err2
	}
	return &result, nil
}
