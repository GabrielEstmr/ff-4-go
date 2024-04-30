package ff_domains

import "reflect"

type Rollout struct {
	key         string
	group       string
	description string
	enabled     bool
	targets     map[string]interface{}
}

func NewTarget(
	key string,
	group string,
	description string,
	enabled bool,
	targets map[string]interface{},
) *Rollout {
	return &Rollout{
		key:         key,
		group:       group,
		description: description,
		enabled:     enabled,
		targets:     targets,
	}
}

func (this Rollout) GetKey() string {
	return this.key
}

func (this Rollout) GetGroup() string {
	return this.group
}

func (this Rollout) GetDescription() string {
	return this.description
}

func (this Rollout) GetEnabled() bool {
	return this.enabled
}

func (this Rollout) GetTargets() map[string]interface{} {
	return this.targets
}

func (this Rollout) GetIsEmpty() bool {
	document := *new(Rollout)
	return reflect.DeepEqual(this, document)
}
