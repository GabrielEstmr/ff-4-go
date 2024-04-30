package ff_domains

import "reflect"

type Feature struct {
	key         string
	group       string
	description string
	value       bool
}

func NewFeature(
	key string,
	group string,
	description string,
	defaultValue bool) *Feature {
	return &Feature{
		key:         key,
		group:       group,
		description: description,
		value:       defaultValue}
}

func (this Feature) IsEmpty() bool {
	document := *new(Feature)
	return reflect.DeepEqual(this, document)
}

func (this Feature) GetKey() string {
	return this.key
}

func (this Feature) GetGroup() string {
	return this.group
}

func (this Feature) GetDescription() string {
	return this.description
}

func (this Feature) GetValue() bool {
	return this.value
}

func (this Feature) SetDefaultValue(defaultValue bool) {
	this.value = defaultValue
}

func (this Feature) IsEnabled() bool {
	return this.value
}

func (this Feature) IsDisabled() bool {
	return !this.IsEnabled()
}

func (this Feature) CloneAsEnabled() Feature {
	return Feature{
		key:         this.key,
		group:       this.group,
		description: this.description,
		value:       true,
	}
}

func (this Feature) CloneAsDisabled() Feature {
	return Feature{
		key:         this.key,
		group:       this.group,
		description: this.description,
		value:       false,
	}
}
