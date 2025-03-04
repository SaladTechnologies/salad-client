/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.7
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerLoggingDatadogTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingDatadogTagsInner{}

// ContainerLoggingDatadogTagsInner struct for ContainerLoggingDatadogTagsInner
type ContainerLoggingDatadogTagsInner struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _ContainerLoggingDatadogTagsInner ContainerLoggingDatadogTagsInner

// NewContainerLoggingDatadogTagsInner instantiates a new ContainerLoggingDatadogTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingDatadogTagsInner(name string, value string) *ContainerLoggingDatadogTagsInner {
	this := ContainerLoggingDatadogTagsInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewContainerLoggingDatadogTagsInnerWithDefaults instantiates a new ContainerLoggingDatadogTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingDatadogTagsInnerWithDefaults() *ContainerLoggingDatadogTagsInner {
	this := ContainerLoggingDatadogTagsInner{}
	return &this
}

// GetName returns the Name field value
func (o *ContainerLoggingDatadogTagsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingDatadogTagsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerLoggingDatadogTagsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *ContainerLoggingDatadogTagsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingDatadogTagsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ContainerLoggingDatadogTagsInner) SetValue(v string) {
	o.Value = v
}

func (o ContainerLoggingDatadogTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingDatadogTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ContainerLoggingDatadogTagsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContainerLoggingDatadogTagsInner := _ContainerLoggingDatadogTagsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerLoggingDatadogTagsInner)

	if err != nil {
		return err
	}

	*o = ContainerLoggingDatadogTagsInner(varContainerLoggingDatadogTagsInner)

	return err
}

type NullableContainerLoggingDatadogTagsInner struct {
	value *ContainerLoggingDatadogTagsInner
	isSet bool
}

func (v NullableContainerLoggingDatadogTagsInner) Get() *ContainerLoggingDatadogTagsInner {
	return v.value
}

func (v *NullableContainerLoggingDatadogTagsInner) Set(val *ContainerLoggingDatadogTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingDatadogTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingDatadogTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingDatadogTagsInner(val *ContainerLoggingDatadogTagsInner) *NullableContainerLoggingDatadogTagsInner {
	return &NullableContainerLoggingDatadogTagsInner{value: val, isSet: true}
}

func (v NullableContainerLoggingDatadogTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingDatadogTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


