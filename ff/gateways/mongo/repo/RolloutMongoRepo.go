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

type RolloutMongoRepo struct {
	_FIELD_KEY                string
	_FIELD_GROUP              string
	_FIELD_DESCRIPTION        string
	_FIELD_ENABLED_ALL        string
	_FIELD_TARGETS            string
	_FIELD_LAST_MODIFIED_DATE string
	ffConfigData              ff.FfClientArgs
}

func NewRolloutMongoRepo(ffConfigData ff.FfClientArgs) *RolloutMongoRepo {
	return &RolloutMongoRepo{
		_FIELD_KEY:                "_id",
		_FIELD_GROUP:              "group",
		_FIELD_DESCRIPTION:        "description",
		_FIELD_ENABLED_ALL:        "enabled_all",
		_FIELD_TARGETS:            "targets",
		_FIELD_LAST_MODIFIED_DATE: "last_modified_date",
		ffConfigData:              ffConfigData,
	}
}

func (this *RolloutMongoRepo) Save(
	rolloutDocument ff_gateways_mongo_documents.RolloutDocument,
) (ff_gateways_mongo_documents.RolloutDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetRolloutColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	rolloutDocument.CreatedDate = now
	rolloutDocument.LastModifiedDate = now

	result, err := collection.InsertOne(context.TODO(), rolloutDocument)
	if err != nil {
		return *new(ff_gateways_mongo_documents.RolloutDocument), err
	}

	key, _ := result.InsertedID.(string)
	rolloutDocument.Key = key
	return rolloutDocument, nil
}

func (this *RolloutMongoRepo) Update(
	rolloutDocument ff_gateways_mongo_documents.RolloutDocument,
) (ff_gateways_mongo_documents.RolloutDocument, error) {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetRolloutColName())
	now := primitive.NewDateTimeFromTime(time.Now())
	rolloutDocument.LastModifiedDate = now

	filter := bson.D{{this._FIELD_KEY, rolloutDocument.Key}}
	update := bson.D{
		{"$set", bson.D{{this._FIELD_GROUP, rolloutDocument.Group}}},
		{"$set", bson.D{{this._FIELD_DESCRIPTION, rolloutDocument.Description}}},
		{"$set", bson.D{{this._FIELD_ENABLED_ALL, rolloutDocument.EnabledAll}}},
		{"$set", bson.D{{this._FIELD_TARGETS, rolloutDocument.Targets}}},
		{"$set", bson.D{{this._FIELD_LAST_MODIFIED_DATE, rolloutDocument.LastModifiedDate}}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return *new(ff_gateways_mongo_documents.RolloutDocument), err
	}
	return rolloutDocument, nil
}

func (this *RolloutMongoRepo) Delete(key string) error {

	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetRolloutColName())
	filter := bson.D{{this._FIELD_KEY, key}}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (this *RolloutMongoRepo) FindById(id string) (*ff_gateways_mongo_documents.RolloutDocument, error) {
	collection := this.ffConfigData.GetDb().Collection(this.ffConfigData.GetRolloutColName())
	var result ff_gateways_mongo_documents.RolloutDocument
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
