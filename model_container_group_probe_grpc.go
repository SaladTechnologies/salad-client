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

// checks if the ContainerGroupProbeGrpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeGrpc{}

// ContainerGroupProbeGrpc Configuration for gRPC-based health probes in container groups, used to determine container health status.
type ContainerGroupProbeGrpc struct {
	// The port number on which the gRPC health check service is exposed.
	Port int32 `json:"port"`
	// The name of the gRPC service that implements the health check protocol.
	Service string `json:"service" validate:"regexp=^.*$"`
}

type _ContainerGroupProbeGrpc ContainerGroupProbeGrpc

// NewContainerGroupProbeGrpc instantiates a new ContainerGroupProbeGrpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeGrpc(port int32, service string) *ContainerGroupProbeGrpc {
	this := ContainerGroupProbeGrpc{}
	this.Port = port
	this.Service = service
	return &this
}

// NewContainerGroupProbeGrpcWithDefaults instantiates a new ContainerGroupProbeGrpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeGrpcWithDefaults() *ContainerGroupProbeGrpc {
	this := ContainerGroupProbeGrpc{}
	return &this
}

// GetPort returns the Port field value
func (o *ContainerGroupProbeGrpc) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeGrpc) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupProbeGrpc) SetPort(v int32) {
	o.Port = v
}

// GetService returns the Service field value
func (o *ContainerGroupProbeGrpc) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeGrpc) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ContainerGroupProbeGrpc) SetService(v string) {
	o.Service = v
}

func (o ContainerGroupProbeGrpc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeGrpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

func (o *ContainerGroupProbeGrpc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"service",
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

	varContainerGroupProbeGrpc := _ContainerGroupProbeGrpc{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupProbeGrpc)

	if err != nil {
		return err
	}

	*o = ContainerGroupProbeGrpc(varContainerGroupProbeGrpc)

	return err
}

type NullableContainerGroupProbeGrpc struct {
	value *ContainerGroupProbeGrpc
	isSet bool
}

func (v NullableContainerGroupProbeGrpc) Get() *ContainerGroupProbeGrpc {
	return v.value
}

func (v *NullableContainerGroupProbeGrpc) Set(val *ContainerGroupProbeGrpc) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeGrpc) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeGrpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeGrpc(val *ContainerGroupProbeGrpc) *NullableContainerGroupProbeGrpc {
	return &NullableContainerGroupProbeGrpc{value: val, isSet: true}
}

func (v NullableContainerGroupProbeGrpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeGrpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


