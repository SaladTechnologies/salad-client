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

// checks if the CreateContainerRegistryAuthenticationGcpGcr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerRegistryAuthenticationGcpGcr{}

// CreateContainerRegistryAuthenticationGcpGcr struct for CreateContainerRegistryAuthenticationGcpGcr
type CreateContainerRegistryAuthenticationGcpGcr struct {
	ServiceKey string `json:"service_key"`
}

type _CreateContainerRegistryAuthenticationGcpGcr CreateContainerRegistryAuthenticationGcpGcr

// NewCreateContainerRegistryAuthenticationGcpGcr instantiates a new CreateContainerRegistryAuthenticationGcpGcr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerRegistryAuthenticationGcpGcr(serviceKey string) *CreateContainerRegistryAuthenticationGcpGcr {
	this := CreateContainerRegistryAuthenticationGcpGcr{}
	this.ServiceKey = serviceKey
	return &this
}

// NewCreateContainerRegistryAuthenticationGcpGcrWithDefaults instantiates a new CreateContainerRegistryAuthenticationGcpGcr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerRegistryAuthenticationGcpGcrWithDefaults() *CreateContainerRegistryAuthenticationGcpGcr {
	this := CreateContainerRegistryAuthenticationGcpGcr{}
	return &this
}

// GetServiceKey returns the ServiceKey field value
func (o *CreateContainerRegistryAuthenticationGcpGcr) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value
// and a boolean to check if the value has been set.
func (o *CreateContainerRegistryAuthenticationGcpGcr) GetServiceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceKey, true
}

// SetServiceKey sets field value
func (o *CreateContainerRegistryAuthenticationGcpGcr) SetServiceKey(v string) {
	o.ServiceKey = v
}

func (o CreateContainerRegistryAuthenticationGcpGcr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerRegistryAuthenticationGcpGcr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service_key"] = o.ServiceKey
	return toSerialize, nil
}

func (o *CreateContainerRegistryAuthenticationGcpGcr) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"service_key",
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

	varCreateContainerRegistryAuthenticationGcpGcr := _CreateContainerRegistryAuthenticationGcpGcr{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateContainerRegistryAuthenticationGcpGcr)

	if err != nil {
		return err
	}

	*o = CreateContainerRegistryAuthenticationGcpGcr(varCreateContainerRegistryAuthenticationGcpGcr)

	return err
}

type NullableCreateContainerRegistryAuthenticationGcpGcr struct {
	value *CreateContainerRegistryAuthenticationGcpGcr
	isSet bool
}

func (v NullableCreateContainerRegistryAuthenticationGcpGcr) Get() *CreateContainerRegistryAuthenticationGcpGcr {
	return v.value
}

func (v *NullableCreateContainerRegistryAuthenticationGcpGcr) Set(val *CreateContainerRegistryAuthenticationGcpGcr) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerRegistryAuthenticationGcpGcr) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerRegistryAuthenticationGcpGcr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerRegistryAuthenticationGcpGcr(val *CreateContainerRegistryAuthenticationGcpGcr) *NullableCreateContainerRegistryAuthenticationGcpGcr {
	return &NullableCreateContainerRegistryAuthenticationGcpGcr{value: val, isSet: true}
}

func (v NullableCreateContainerRegistryAuthenticationGcpGcr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerRegistryAuthenticationGcpGcr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


