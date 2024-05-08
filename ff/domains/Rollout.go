package ff_domains

import (
	"reflect"
	"time"
)

// Rollout is the resource for multiple-values temporary features to control processes rollouts.
type Rollout struct {
	key              string
	group            string
	description      string
	enabledAll       bool
	targets          map[string]interface{}
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewRolloutAllArgs(
	key string,
	group string,
	description string,
	enabledAll bool,
	targets map[string]interface{},
	createdDate time.Time,
	lastModifiedDate time.Time,
) *Rollout {
	return &Rollout{
		key:              key,
		group:            group,
		description:      description,
		enabledAll:       enabledAll,
		targets:          targets,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func NewRollout(
	key string,
	group string,
	description string,
	enabledAll bool,
	targets map[string]interface{},
) *Rollout {
	return &Rollout{
		key:         key,
		group:       group,
		description: description,
		enabledAll:  enabledAll,
		targets:     targets,
	}
}

func (this *Rollout) GetKey() string {
	return this.key
}

func (this *Rollout) GetGroup() string {
	return this.group
}

func (this *Rollout) GetDescription() string {
	return this.description
}

func (this *Rollout) GetEnabledAll() bool {
	return this.enabledAll
}

func (this *Rollout) GetTargets() map[string]interface{} {
	return this.targets
}

func (this *Rollout) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this *Rollout) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this *Rollout) IsEmpty() bool {
	document := new(Rollout)
	return reflect.DeepEqual(this, document)
}

func (this *Rollout) RemoveTarget(target string) {
	delete(this.targets, target)
}

func (this *Rollout) AddTarget(target string) {
	this.AddTargets(map[string]interface{}{
		target: target,
	})
}

func (this *Rollout) AddTargets(targets map[string]interface{}) {
	targetsToUpdate := this.targets
	for key, value := range targets {
		targetsToUpdate[key] = value
	}
	this.targets = targetsToUpdate
}

func (this *Rollout) SetEnabledAll(enabledAll bool) {
	this.enabledAll = enabledAll
}

func (this *Rollout) IsTargetInRollout(target string) bool {
	_, ok := this.targets[target]
	return ok
}

func (this *Rollout) IsEnabledAllOrTargetInRollout(target string) bool {
	if this.enabledAll {
		return true
	}
	return this.IsTargetInRollout(target)
}
