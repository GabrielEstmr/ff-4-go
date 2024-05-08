package ff_gateways_mongo_documents

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type RolloutDocument struct {
	Key              string                 `json:"_id,omitempty" bson:"_id,omitempty"`
	Group            string                 `json:"group,omitempty" bson:"group,omitempty"`
	Description      string                 `json:"description,omitempty" bson:"description,omitempty"`
	EnabledAll       bool                   `json:"enabled_all" bson:"enabled"`
	Targets          map[string]interface{} `json:"targets,omitempty" bson:"targets,omitempty"`
	CreatedDate      primitive.DateTime     `json:"created_date,omitempty" bson:"created_date"`
	LastModifiedDate primitive.DateTime     `json:"last_modified_date,omitempty" bson:"last_modified_date"`
}

func NewRolloutDocument(
	rollout ff_domains.Rollout,
) *RolloutDocument {
	return &RolloutDocument{
		Key:              rollout.GetKey(),
		Group:            rollout.GetGroup(),
		Description:      rollout.GetDescription(),
		EnabledAll:       rollout.GetEnabledAll(),
		Targets:          rollout.GetTargets(),
		CreatedDate:      primitive.NewDateTimeFromTime(rollout.GetCreatedDate()),
		LastModifiedDate: primitive.NewDateTimeFromTime(rollout.GetLastModifiedDate()),
	}
}

func (this RolloutDocument) ToDomain() ff_domains.Rollout {
	if this.IsEmpty() {
		return *new(ff_domains.Rollout)
	}
	return *ff_domains.NewRolloutAllArgs(
		this.Key,
		this.Group,
		this.Description,
		this.EnabledAll,
		this.Targets,
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}

func (this RolloutDocument) IsEmpty() bool {
	document := *new(RolloutDocument)
	return reflect.DeepEqual(this, document)
}
