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

// checks if the ContainerGroupProbeHttp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbeHttp{}

// ContainerGroupProbeHttp Defines HTTP probe configuration for container health checks within a container group.
type ContainerGroupProbeHttp struct {
	// A collection of HTTP header name-value pairs used for configuring requests and responses in container group endpoints. Each header consists of a name and its corresponding value.
	Headers []ContainerGroupProbeHttpHeader `json:"headers"`
	// The HTTP path that will be probed to check container health.
	Path string `json:"path" validate:"regexp=^.*$"`
	// The TCP port number to which the HTTP request will be sent.
	Port int32 `json:"port"`
	Scheme NullableContainerProbeHttpScheme `json:"scheme"`
}

type _ContainerGroupProbeHttp ContainerGroupProbeHttp

// NewContainerGroupProbeHttp instantiates a new ContainerGroupProbeHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbeHttp(headers []ContainerGroupProbeHttpHeader, path string, port int32, scheme NullableContainerProbeHttpScheme) *ContainerGroupProbeHttp {
	this := ContainerGroupProbeHttp{}
	this.Headers = headers
	this.Path = path
	this.Port = port
	this.Scheme = scheme
	return &this
}

// NewContainerGroupProbeHttpWithDefaults instantiates a new ContainerGroupProbeHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeHttpWithDefaults() *ContainerGroupProbeHttp {
	this := ContainerGroupProbeHttp{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *ContainerGroupProbeHttp) GetHeaders() []ContainerGroupProbeHttpHeader {
	if o == nil {
		var ret []ContainerGroupProbeHttpHeader
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetHeadersOk() ([]ContainerGroupProbeHttpHeader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *ContainerGroupProbeHttp) SetHeaders(v []ContainerGroupProbeHttpHeader) {
	o.Headers = v
}

// GetPath returns the Path field value
func (o *ContainerGroupProbeHttp) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ContainerGroupProbeHttp) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *ContainerGroupProbeHttp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbeHttp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerGroupProbeHttp) SetPort(v int32) {
	o.Port = v
}

// GetScheme returns the Scheme field value
// If the value is explicit nil, the zero value for ContainerProbeHttpScheme will be returned
func (o *ContainerGroupProbeHttp) GetScheme() ContainerProbeHttpScheme {
	if o == nil || o.Scheme.Get() == nil {
		var ret ContainerProbeHttpScheme
		return ret
	}

	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerGroupProbeHttp) GetSchemeOk() (*ContainerProbeHttpScheme, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// SetScheme sets field value
func (o *ContainerGroupProbeHttp) SetScheme(v ContainerProbeHttpScheme) {
	o.Scheme.Set(&v)
}

func (o ContainerGroupProbeHttp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbeHttp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers
	toSerialize["path"] = o.Path
	toSerialize["port"] = o.Port
	toSerialize["scheme"] = o.Scheme.Get()
	return toSerialize, nil
}

func (o *ContainerGroupProbeHttp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headers",
		"path",
		"port",
		"scheme",
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

	varContainerGroupProbeHttp := _ContainerGroupProbeHttp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupProbeHttp)

	if err != nil {
		return err
	}

	*o = ContainerGroupProbeHttp(varContainerGroupProbeHttp)

	return err
}

type NullableContainerGroupProbeHttp struct {
	value *ContainerGroupProbeHttp
	isSet bool
}

func (v NullableContainerGroupProbeHttp) Get() *ContainerGroupProbeHttp {
	return v.value
}

func (v *NullableContainerGroupProbeHttp) Set(val *ContainerGroupProbeHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbeHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbeHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbeHttp(val *ContainerGroupProbeHttp) *NullableContainerGroupProbeHttp {
	return &NullableContainerGroupProbeHttp{value: val, isSet: true}
}

func (v NullableContainerGroupProbeHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbeHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


