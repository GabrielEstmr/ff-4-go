package ff_domains

import (
	"reflect"
	"time"
)

// FeatureFlag is the resource for boolean features.
type FeatureFlag struct {
	key              string
	group            string
	description      string
	value            bool
	createdDate      time.Time
	lastModifiedDate time.Time
}

func NewFeatureFlagAllArgs(
	key string,
	group string,
	description string,
	value bool,
	createdDate time.Time,
	lastModifiedDate time.Time,
) *FeatureFlag {
	return &FeatureFlag{
		key:              key,
		group:            group,
		description:      description,
		value:            value,
		createdDate:      createdDate,
		lastModifiedDate: lastModifiedDate,
	}
}

func NewFeatureFlag(
	key string,
	group string,
	description string,
	value bool,
) *FeatureFlag {
	return &FeatureFlag{
		key:         key,
		group:       group,
		description: description,
		value:       value,
	}
}

func (this FeatureFlag) IsEmpty() bool {
	document := *new(FeatureFlag)
	return reflect.DeepEqual(this, document)
}

func (this FeatureFlag) GetKey() string {
	return this.key
}

func (this FeatureFlag) GetGroup() string {
	return this.group
}

func (this FeatureFlag) GetDescription() string {
	return this.description
}

func (this FeatureFlag) GetValue() bool {
	return this.value
}

func (this FeatureFlag) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this FeatureFlag) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this FeatureFlag) SetDefaultValue(defaultValue bool) {
	this.value = defaultValue
}

func (this FeatureFlag) IsEnabled() bool {
	return this.value
}

func (this FeatureFlag) IsDisabled() bool {
	return !this.IsEnabled()
}

func (this FeatureFlag) CloneAsEnabled() FeatureFlag {
	return FeatureFlag{
		key:         this.key,
		group:       this.group,
		description: this.description,
		value:       true,
	}
}

func (this FeatureFlag) CloneAsDisabled() FeatureFlag {
	return FeatureFlag{
		key:         this.key,
		group:       this.group,
		description: this.description,
		value:       false,
	}
}
