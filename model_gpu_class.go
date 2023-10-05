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

// checks if the GpuClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuClass{}

// GpuClass Represents a GPU Class
type GpuClass struct {
	// The GPU class id
	Id string `json:"id"`
	// The GPU class name
	Name string `json:"name"`
	IsHighDemand *bool `json:"is_high_demand,omitempty"`
}

// NewGpuClass instantiates a new GpuClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuClass(id string, name string) *GpuClass {
	this := GpuClass{}
	this.Id = id
	this.Name = name
	return &this
}

// NewGpuClassWithDefaults instantiates a new GpuClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuClassWithDefaults() *GpuClass {
	this := GpuClass{}
	return &this
}

// GetId returns the Id field value
func (o *GpuClass) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GpuClass) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GpuClass) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GpuClass) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GpuClass) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GpuClass) SetName(v string) {
	o.Name = v
}

// GetIsHighDemand returns the IsHighDemand field value if set, zero value otherwise.
func (o *GpuClass) GetIsHighDemand() bool {
	if o == nil || IsNil(o.IsHighDemand) {
		var ret bool
		return ret
	}
	return *o.IsHighDemand
}

// GetIsHighDemandOk returns a tuple with the IsHighDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GpuClass) GetIsHighDemandOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHighDemand) {
		return nil, false
	}
	return o.IsHighDemand, true
}

// HasIsHighDemand returns a boolean if a field has been set.
func (o *GpuClass) HasIsHighDemand() bool {
	if o != nil && !IsNil(o.IsHighDemand) {
		return true
	}

	return false
}

// SetIsHighDemand gets a reference to the given bool and assigns it to the IsHighDemand field.
func (o *GpuClass) SetIsHighDemand(v bool) {
	o.IsHighDemand = &v
}

func (o GpuClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.IsHighDemand) {
		toSerialize["is_high_demand"] = o.IsHighDemand
	}
	return toSerialize, nil
}

type NullableGpuClass struct {
	value *GpuClass
	isSet bool
}

func (v NullableGpuClass) Get() *GpuClass {
	return v.value
}

func (v *NullableGpuClass) Set(val *GpuClass) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuClass) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuClass(val *GpuClass) *NullableGpuClass {
	return &NullableGpuClass{value: val, isSet: true}
}

func (v NullableGpuClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


