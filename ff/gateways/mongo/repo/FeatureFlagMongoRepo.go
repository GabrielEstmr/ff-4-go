package ff_gateways_mongo_repo

import (
	"context"
	"errors"
	ff_gateways_mongo_documents "github.com/GabrielEstmr/ff-4-go/ff/gateways/mongo/documents"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type FeatureFlagMongoRepo struct {
	_FIELD_KEY   string
	ffConfigData ff.FfClientArgs
}

func NewFeatureFlagMongoRepo(ffConfigData ff.FfClientArgs) *FeatureFlagMongoRepo {
	return &FeatureFlagMongoRepo{
		_FIELD_KEY:   "_id",
		ffConfigData: ffConfigData,
	}
}

func (this *FeatureFlagMongoRepo) Save(
	featureFlagDocument ff_gateways_mongo_documents.FeatureFlagDocument,
) (ff_gateways_mongo_documents.FeatureFlagDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesFlagColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	featureFlagDocument.CreatedDate = now
	featureFlagDocument.LastModifiedDate = now

	result, err := collection.InsertOne(context.TODO(), featureFlagDocument)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeatureFlagDocument), err
	}

	key, _ := result.InsertedID.(string)
	featureFlagDocument.Key = key
	return featureFlagDocument, nil
}

func (this *FeatureFlagMongoRepo) Update(
	featureFlagDocument ff_gateways_mongo_documents.FeatureFlagDocument,
) (ff_gateways_mongo_documents.FeatureFlagDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesFlagColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	featureFlagDocument.LastModifiedDate = now

	filter := bson.D{{this._FIELD_KEY, featureFlagDocument.Key}}
	update := bson.D{{"$set", bson.D{{"value", featureFlagDocument.Value}}}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(ff_gateways_mongo_documents.FeatureFlagDocument), err
	}
	return featureFlagDocument, nil
}

func (this *FeatureFlagMongoRepo) Delete(key string) error {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesFlagColName())
	filter := bson.D{{this._FIELD_KEY, key}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (this *FeatureFlagMongoRepo) FindById(id string) (*ff_gateways_mongo_documents.FeatureFlagDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetFeaturesFlagColName())
	var result ff_gateways_mongo_documents.FeatureFlagDocument
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
