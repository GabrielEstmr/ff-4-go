package ff_mongo_repo

import (
	"context"
	"errors"
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/mongo/documents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const _KEY = "_id"

type FeaturesMongoRepo struct {
	ffConfigData *ff.FfConfigData
}

func NewFeaturesMongoRepo(ffConfigData *ff.FfConfigData) *FeaturesMongoRepo {
	return &FeaturesMongoRepo{ffConfigData: ffConfigData}
}

func (this *FeaturesMongoRepo) Save(
	feature ff_mongo_documents.FeaturesDataDocument,
) (ff_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())

	result, err := collection.InsertOne(context.TODO(), feature)
	if err != nil {
		return *new(ff_mongo_documents.FeaturesDataDocument), err
	}

	key, _ := result.InsertedID.(string)
	feature.Key = key
	return feature, nil
}

func (this *FeaturesMongoRepo) Update(
	feature ff_mongo_documents.FeaturesDataDocument,
) (ff_mongo_documents.FeaturesDataDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())
	filter := bson.D{{_KEY, feature.Key}}
	update := bson.D{{"$set", bson.D{{"defaultValue", feature.DefaultValue}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(ff_mongo_documents.FeaturesDataDocument), err
	}
	return feature, nil
}

func (this *FeaturesMongoRepo) FindById(id string) (*ff_mongo_documents.FeaturesDataDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesDbName())
	var result ff_mongo_documents.FeaturesDataDocument
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
