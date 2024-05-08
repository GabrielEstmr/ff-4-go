package ff_gateways_ws_resources

import (
	ff_domains "baseapplicationgo/main/configs/ff/lib/domains"
	"reflect"
	"time"
)

type RolloutResponse struct {
	Key              string                 `json:"key,omitempty"`
	Group            string                 `json:"group,omitempty"`
	Description      string                 `json:"description,omitempty"`
	EnabledAll       bool                   `json:"enabled_all"`
	Targets          map[string]interface{} `json:"targets,omitempty"`
	CreatedDate      time.Time              `json:"created_date,omitempty"`
	LastModifiedDate time.Time              `json:"last_modified_date,omitempty"`
}

func NewRolloutResponse(rollout ff_domains.Rollout) *RolloutResponse {
	return &RolloutResponse{
		Key:              rollout.GetKey(),
		Group:            rollout.GetGroup(),
		Description:      rollout.GetDescription(),
		EnabledAll:       rollout.GetEnabledAll(),
		Targets:          rollout.GetTargets(),
		CreatedDate:      rollout.GetCreatedDate(),
		LastModifiedDate: rollout.GetLastModifiedDate(),
	}
}

func (this RolloutResponse) IsEmpty() bool {
	document := *new(RolloutResponse)
	return reflect.DeepEqual(this, document)
}
