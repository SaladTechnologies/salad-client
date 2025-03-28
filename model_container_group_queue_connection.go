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

// checks if the ContainerGroupQueueConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupQueueConnection{}

// ContainerGroupQueueConnection Configuration for connecting a container group to a message queue system, enabling asynchronous communication between services.
type ContainerGroupQueueConnection struct {
	// The endpoint path for accessing the queue service, relative to the base URL of the queue server.
	Path string `json:"path" validate:"regexp=^.*$"`
	// The network port number used to connect to the queue service. Must be a valid TCP/IP port between 1 and 65535.
	Port int32 `json:"port"`
	// Unique identifier for the queue. Must start with a lowercase letter, can contain lowercase letters, numbers, and hyphens, and must end with a letter or number.
	QueueName string `json:"queue_name" validate:"regexp=^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
}

type _ContainerGroupQueueConnection ContainerGroupQueueConnection

// NewContainerGroupQueueConnection instantiates a new ContainerGroupQueueConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupQueueConnection(path string, port int32, queueName string) *ContainerGroupQueueConnection {
	this := ContainerGroupQueueConnection{}
	this.Path = path
	this.Port = port
	this.QueueName = queueName
	return &this
}

// NewContainerGroupQueueConnectionWithDefaults instantiates a new ContainerGroupQueueConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupQueueConnectionWithDefaults() *ContainerGroupQueueConnection {
	this := ContainerGroupQueueConnection{}
	return &this
}

// GetPath returns the Path field value
func (o *ContainerGroupQueueConnection) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueConnection) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ContainerGroupQueueConnection) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *ContainerGroupQueueConnection) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueConnection) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupQueueConnection) SetPort(v int32) {
	o.Port = v
}

// GetQueueName returns the QueueName field value
func (o *ContainerGroupQueueConnection) GetQueueName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupQueueConnection) GetQueueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueName, true
}

// SetQueueName sets field value
func (o *ContainerGroupQueueConnection) SetQueueName(v string) {
	o.QueueName = v
}

func (o ContainerGroupQueueConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupQueueConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	toSerialize["queue_name"] = o.QueueName
	return toSerialize, nil
}

func (o *ContainerGroupQueueConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"port",
		"queue_name",
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

	varContainerGroupQueueConnection := _ContainerGroupQueueConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupQueueConnection)

	if err != nil {
		return err
	}

	*o = ContainerGroupQueueConnection(varContainerGroupQueueConnection)

	return err
}

type NullableContainerGroupQueueConnection struct {
	value *ContainerGroupQueueConnection
	isSet bool
}

func (v NullableContainerGroupQueueConnection) Get() *ContainerGroupQueueConnection {
	return v.value
}

func (v *NullableContainerGroupQueueConnection) Set(val *ContainerGroupQueueConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupQueueConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupQueueConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupQueueConnection(val *ContainerGroupQueueConnection) *NullableContainerGroupQueueConnection {
	return &NullableContainerGroupQueueConnection{value: val, isSet: true}
}

func (v NullableContainerGroupQueueConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupQueueConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


