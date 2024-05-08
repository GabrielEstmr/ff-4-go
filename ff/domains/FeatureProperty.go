package ff_domains

import (
	"reflect"
	"time"
)

// FeatureProperty is the resource for multiple-values features.
type FeatureProperty struct {
	key              string
	group            string
	description      string
	enabled          bool
	values           map[string]interface{}
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewFeaturePropertyAllArgs(
	key string,
	group string,
	description string,
	enabled bool,
	values map[string]interface{},
	createdDate time.Time,
	lastModifiedDate time.Time,
) *FeatureProperty {
	return &FeatureProperty{
		key:              key,
		group:            group,
		description:      description,
		enabled:          enabled,
		values:           values,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func NewFeatureProperty(
	key string,
	group string,
	description string,
	enabled bool,
	values map[string]interface{},
) *FeatureProperty {
	return &FeatureProperty{
		key:         key,
		group:       group,
		description: description,
		enabled:     enabled,
		values:      values,
	}
}

func (this *FeatureProperty) SetEnabled(enabled bool) {
	this.enabled = enabled
}

func (this *FeatureProperty) GetKey() string {
	return this.key
}

func (this *FeatureProperty) GetGroup() string {
	return this.group
}

func (this *FeatureProperty) GetDescription() string {
	return this.description
}

func (this *FeatureProperty) GetEnabled() bool {
	return this.enabled
}

func (this *FeatureProperty) GetValues() map[string]interface{} {
	return this.values
}

func (this *FeatureProperty) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this *FeatureProperty) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this *FeatureProperty) IsEmpty() bool {
	document := new(FeatureProperty)
	return reflect.DeepEqual(this, document)
}

func (this *FeatureProperty) RemoveValue(value string) {
	if this.values == nil {
		return
	}
	delete(this.values, value)
}

func (this *FeatureProperty) AddValue(value string) {
	this.AddValues(map[string]interface{}{
		value: value,
	})
}

func (this *FeatureProperty) AddValues(values map[string]interface{}) {
	if this.values == nil {
		this.values = make(map[string]interface{})
	}
	for k, v := range values {
		this.values[k] = v
	}
}

func (this *FeatureProperty) IsValueInProperty(target string) bool {
	_, ok := this.values[target]
	return ok
}
