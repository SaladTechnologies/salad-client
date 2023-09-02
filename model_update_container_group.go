/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateContainerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContainerGroup{}

// UpdateContainerGroup Represents a request to update a container group
type UpdateContainerGroup struct {
	DisplayName NullableString `json:"display_name,omitempty"`
	Replicas NullableInt32 `json:"replicas,omitempty"`
}

// NewUpdateContainerGroup instantiates a new UpdateContainerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContainerGroup() *UpdateContainerGroup {
	this := UpdateContainerGroup{}
	return &this
}

// NewUpdateContainerGroupWithDefaults instantiates a new UpdateContainerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContainerGroupWithDefaults() *UpdateContainerGroup {
	this := UpdateContainerGroup{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainerGroup) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainerGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UpdateContainerGroup) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UpdateContainerGroup) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UpdateContainerGroup) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainerGroup) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas.Get()) {
		var ret int32
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainerGroup) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasReplicas() bool {
	if o != nil && o.Replicas.IsSet() {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given NullableInt32 and assigns it to the Replicas field.
func (o *UpdateContainerGroup) SetReplicas(v int32) {
	o.Replicas.Set(&v)
}
// SetReplicasNil sets the value for Replicas to be an explicit nil
func (o *UpdateContainerGroup) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
func (o *UpdateContainerGroup) UnsetReplicas() {
	o.Replicas.Unset()
}

func (o UpdateContainerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContainerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	return toSerialize, nil
}

type NullableUpdateContainerGroup struct {
	value *UpdateContainerGroup
	isSet bool
}

func (v NullableUpdateContainerGroup) Get() *UpdateContainerGroup {
	return v.value
}

func (v *NullableUpdateContainerGroup) Set(val *UpdateContainerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContainerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContainerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContainerGroup(val *UpdateContainerGroup) *NullableUpdateContainerGroup {
	return &NullableUpdateContainerGroup{value: val, isSet: true}
}

func (v NullableUpdateContainerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContainerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


