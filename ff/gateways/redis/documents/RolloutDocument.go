package ff_gateways_redis_documents

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
	"time"
)

type RolloutDocument struct {
	Key              string                 `json:"_id,omitempty"`
	Group            string                 `json:"group,omitempty"`
	Description      string                 `json:"description,omitempty"`
	EnabledAll       bool                   `json:"enabled_all"`
	Targets          map[string]interface{} `json:"targets,omitempty"`
	CreatedDate      time.Time              `json:"created_date,omitempty"`
	LastModifiedDate time.Time              `json:"last_modified_date,omitempty"`
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
		CreatedDate:      rollout.GetCreatedDate(),
		LastModifiedDate: rollout.GetLastModifiedDate(),
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
		this.CreatedDate,
		this.LastModifiedDate,
	)
}

func (this RolloutDocument) IsEmpty() bool {
	document := *new(RolloutDocument)
	return reflect.DeepEqual(this, document)
}
