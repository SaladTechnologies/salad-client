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

// checks if the CreateContainerGroupNetworking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerGroupNetworking{}

// CreateContainerGroupNetworking Represents container group networking parameters
type CreateContainerGroupNetworking struct {
	Protocol ContainerNetworkingProtocol `json:"protocol"`
	Port int32 `json:"port"`
	Auth bool `json:"auth"`
}

// NewCreateContainerGroupNetworking instantiates a new CreateContainerGroupNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerGroupNetworking(protocol ContainerNetworkingProtocol, port int32, auth bool) *CreateContainerGroupNetworking {
	this := CreateContainerGroupNetworking{}
	this.Protocol = protocol
	this.Port = port
	this.Auth = auth
	return &this
}

// NewCreateContainerGroupNetworkingWithDefaults instantiates a new CreateContainerGroupNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerGroupNetworkingWithDefaults() *CreateContainerGroupNetworking {
	this := CreateContainerGroupNetworking{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *CreateContainerGroupNetworking) GetProtocol() ContainerNetworkingProtocol {
	if o == nil {
		var ret ContainerNetworkingProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroupNetworking) GetProtocolOk() (*ContainerNetworkingProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateContainerGroupNetworking) SetProtocol(v ContainerNetworkingProtocol) {
	o.Protocol = v
}

// GetPort returns the Port field value
func (o *CreateContainerGroupNetworking) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroupNetworking) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateContainerGroupNetworking) SetPort(v int32) {
	o.Port = v
}

// GetAuth returns the Auth field value
func (o *CreateContainerGroupNetworking) GetAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroupNetworking) GetAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *CreateContainerGroupNetworking) SetAuth(v bool) {
	o.Auth = v
}

func (o CreateContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerGroupNetworking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["port"] = o.Port
	toSerialize["auth"] = o.Auth
	return toSerialize, nil
}

type NullableCreateContainerGroupNetworking struct {
	value *CreateContainerGroupNetworking
	isSet bool
}

func (v NullableCreateContainerGroupNetworking) Get() *CreateContainerGroupNetworking {
	return v.value
}

func (v *NullableCreateContainerGroupNetworking) Set(val *CreateContainerGroupNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerGroupNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerGroupNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerGroupNetworking(val *CreateContainerGroupNetworking) *NullableCreateContainerGroupNetworking {
	return &NullableCreateContainerGroupNetworking{value: val, isSet: true}
}

func (v NullableCreateContainerGroupNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerGroupNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


