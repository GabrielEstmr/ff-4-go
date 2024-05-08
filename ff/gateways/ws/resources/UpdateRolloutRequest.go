package ff_gateways_ws_resources

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	"reflect"
)

type UpdateRolloutRequest struct {
	Group       string                 `json:"group,omitempty"`
	Description string                 `json:"description,omitempty"`
	EnabledAll  bool                   `json:"enabled_all,omitempty"`
	Targets     map[string]interface{} `json:"targets,omitempty"`
}

func (this UpdateRolloutRequest) IsEmpty() bool {
	document := *new(UpdateRolloutRequest)
	return reflect.DeepEqual(this, document)
}

func (this UpdateRolloutRequest) ToDomain(key string) ff_domains.Rollout {
	return *ff_domains.NewRollout(
		key,
		this.Group,
		this.Description,
		this.EnabledAll,
		this.Targets,
	)
}
