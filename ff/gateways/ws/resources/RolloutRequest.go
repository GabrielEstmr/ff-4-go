package ff_gateways_ws_resources

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
)

type RolloutRequest struct {
	Key         string                 `json:"key,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	EnabledAll  bool                   `json:"enabled_all,omitempty"`
	Targets     map[string]interface{} `json:"targets,omitempty"`
}

func (this RolloutRequest) IsEmpty() bool {
	document := *new(RolloutRequest)
	return reflect.DeepEqual(this, document)
}

func (this RolloutRequest) ToDomain() ff_domains.Rollout {
	return *ff_domains.NewRollout(
		this.Key,
		this.Group,
		this.Description,
		this.EnabledAll,
		this.Targets,
	)
}
