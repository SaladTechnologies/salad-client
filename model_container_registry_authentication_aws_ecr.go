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

// checks if the ContainerRegistryAuthenticationAwsEcr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerRegistryAuthenticationAwsEcr{}

// ContainerRegistryAuthenticationAwsEcr Authentication details for AWS Elastic Container Registry (ECR)
type ContainerRegistryAuthenticationAwsEcr struct {
	// AWS access key ID used for ECR authentication
	AccessKeyId string `json:"access_key_id" validate:"regexp=^.*$"`
	// AWS secret access key used for ECR authentication
	SecretAccessKey string `json:"secret_access_key" validate:"regexp=^.*$"`
}

type _ContainerRegistryAuthenticationAwsEcr ContainerRegistryAuthenticationAwsEcr

// NewContainerRegistryAuthenticationAwsEcr instantiates a new ContainerRegistryAuthenticationAwsEcr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRegistryAuthenticationAwsEcr(accessKeyId string, secretAccessKey string) *ContainerRegistryAuthenticationAwsEcr {
	this := ContainerRegistryAuthenticationAwsEcr{}
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewContainerRegistryAuthenticationAwsEcrWithDefaults instantiates a new ContainerRegistryAuthenticationAwsEcr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRegistryAuthenticationAwsEcrWithDefaults() *ContainerRegistryAuthenticationAwsEcr {
	this := ContainerRegistryAuthenticationAwsEcr{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *ContainerRegistryAuthenticationAwsEcr) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryAuthenticationAwsEcr) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *ContainerRegistryAuthenticationAwsEcr) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *ContainerRegistryAuthenticationAwsEcr) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *ContainerRegistryAuthenticationAwsEcr) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *ContainerRegistryAuthenticationAwsEcr) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o ContainerRegistryAuthenticationAwsEcr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerRegistryAuthenticationAwsEcr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_key_id"] = o.AccessKeyId
	toSerialize["secret_access_key"] = o.SecretAccessKey
	return toSerialize, nil
}

func (o *ContainerRegistryAuthenticationAwsEcr) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_key_id",
		"secret_access_key",
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

	varContainerRegistryAuthenticationAwsEcr := _ContainerRegistryAuthenticationAwsEcr{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerRegistryAuthenticationAwsEcr)

	if err != nil {
		return err
	}

	*o = ContainerRegistryAuthenticationAwsEcr(varContainerRegistryAuthenticationAwsEcr)

	return err
}

type NullableContainerRegistryAuthenticationAwsEcr struct {
	value *ContainerRegistryAuthenticationAwsEcr
	isSet bool
}

func (v NullableContainerRegistryAuthenticationAwsEcr) Get() *ContainerRegistryAuthenticationAwsEcr {
	return v.value
}

func (v *NullableContainerRegistryAuthenticationAwsEcr) Set(val *ContainerRegistryAuthenticationAwsEcr) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRegistryAuthenticationAwsEcr) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRegistryAuthenticationAwsEcr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRegistryAuthenticationAwsEcr(val *ContainerRegistryAuthenticationAwsEcr) *NullableContainerRegistryAuthenticationAwsEcr {
	return &NullableContainerRegistryAuthenticationAwsEcr{value: val, isSet: true}
}

func (v NullableContainerRegistryAuthenticationAwsEcr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRegistryAuthenticationAwsEcr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


