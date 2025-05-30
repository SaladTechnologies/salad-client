/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.11
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ContainerGroupInstanceCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupInstanceCollection{}

// ContainerGroupInstanceCollection A collection of container group instances returned as part of a paginated response or batch operation result.
type ContainerGroupInstanceCollection struct {
	// An array of container group instances, each representing a deployed container group with its current state and configuration information.
	Instances []ContainerGroupInstance `json:"instances"`
}

type _ContainerGroupInstanceCollection ContainerGroupInstanceCollection

// NewContainerGroupInstanceCollection instantiates a new ContainerGroupInstanceCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupInstanceCollection(instances []ContainerGroupInstance) *ContainerGroupInstanceCollection {
	this := ContainerGroupInstanceCollection{}
	this.Instances = instances
	return &this
}

// NewContainerGroupInstanceCollectionWithDefaults instantiates a new ContainerGroupInstanceCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupInstanceCollectionWithDefaults() *ContainerGroupInstanceCollection {
	this := ContainerGroupInstanceCollection{}
	return &this
}

// GetInstances returns the Instances field value
func (o *ContainerGroupInstanceCollection) GetInstances() []ContainerGroupInstance {
	if o == nil {
		var ret []ContainerGroupInstance
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupInstanceCollection) GetInstancesOk() ([]ContainerGroupInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *ContainerGroupInstanceCollection) SetInstances(v []ContainerGroupInstance) {
	o.Instances = v
}

func (o ContainerGroupInstanceCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupInstanceCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
	return toSerialize, nil
}

func (o *ContainerGroupInstanceCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instances",
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

	varContainerGroupInstanceCollection := _ContainerGroupInstanceCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupInstanceCollection)

	if err != nil {
		return err
	}

	*o = ContainerGroupInstanceCollection(varContainerGroupInstanceCollection)

	return err
}

type NullableContainerGroupInstanceCollection struct {
	value *ContainerGroupInstanceCollection
	isSet bool
}

func (v NullableContainerGroupInstanceCollection) Get() *ContainerGroupInstanceCollection {
	return v.value
}

func (v *NullableContainerGroupInstanceCollection) Set(val *ContainerGroupInstanceCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupInstanceCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupInstanceCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupInstanceCollection(val *ContainerGroupInstanceCollection) *NullableContainerGroupInstanceCollection {
	return &NullableContainerGroupInstanceCollection{value: val, isSet: true}
}

func (v NullableContainerGroupInstanceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupInstanceCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


