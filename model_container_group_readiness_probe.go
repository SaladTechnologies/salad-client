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

// checks if the ContainerGroupReadinessProbe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupReadinessProbe{}

// ContainerGroupReadinessProbe Represents the container group readiness probe
type ContainerGroupReadinessProbe struct {
	Tcp *ContainerGroupProbeTcp `json:"tcp,omitempty"`
	Http *ContainerGroupProbeHttp `json:"http,omitempty"`
	Grpc *ContainerGroupProbeGrpc `json:"grpc,omitempty"`
	Exec *ContainerGroupProbeExec `json:"exec,omitempty"`
	InitialDelaySeconds int32 `json:"initial_delay_seconds"`
	PeriodSeconds int32 `json:"period_seconds"`
	TimeoutSeconds int32 `json:"timeout_seconds"`
	SuccessThreshold int32 `json:"success_threshold"`
	FailureThreshold int32 `json:"failure_threshold"`
}

type _ContainerGroupReadinessProbe ContainerGroupReadinessProbe

// NewContainerGroupReadinessProbe instantiates a new ContainerGroupReadinessProbe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupReadinessProbe(initialDelaySeconds int32, periodSeconds int32, timeoutSeconds int32, successThreshold int32, failureThreshold int32) *ContainerGroupReadinessProbe {
	this := ContainerGroupReadinessProbe{}
	this.InitialDelaySeconds = initialDelaySeconds
	this.PeriodSeconds = periodSeconds
	this.TimeoutSeconds = timeoutSeconds
	this.SuccessThreshold = successThreshold
	this.FailureThreshold = failureThreshold
	return &this
}

// NewContainerGroupReadinessProbeWithDefaults instantiates a new ContainerGroupReadinessProbe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupReadinessProbeWithDefaults() *ContainerGroupReadinessProbe {
	this := ContainerGroupReadinessProbe{}
	var initialDelaySeconds int32 = 0
	this.InitialDelaySeconds = initialDelaySeconds
	var periodSeconds int32 = 1
	this.PeriodSeconds = periodSeconds
	var timeoutSeconds int32 = 1
	this.TimeoutSeconds = timeoutSeconds
	var successThreshold int32 = 1
	this.SuccessThreshold = successThreshold
	var failureThreshold int32 = 3
	this.FailureThreshold = failureThreshold
	return &this
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *ContainerGroupReadinessProbe) GetTcp() ContainerGroupProbeTcp {
	if o == nil || IsNil(o.Tcp) {
		var ret ContainerGroupProbeTcp
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetTcpOk() (*ContainerGroupProbeTcp, bool) {
	if o == nil || IsNil(o.Tcp) {
		return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *ContainerGroupReadinessProbe) HasTcp() bool {
	if o != nil && !IsNil(o.Tcp) {
		return true
	}

	return false
}

// SetTcp gets a reference to the given ContainerGroupProbeTcp and assigns it to the Tcp field.
func (o *ContainerGroupReadinessProbe) SetTcp(v ContainerGroupProbeTcp) {
	o.Tcp = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *ContainerGroupReadinessProbe) GetHttp() ContainerGroupProbeHttp {
	if o == nil || IsNil(o.Http) {
		var ret ContainerGroupProbeHttp
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetHttpOk() (*ContainerGroupProbeHttp, bool) {
	if o == nil || IsNil(o.Http) {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *ContainerGroupReadinessProbe) HasHttp() bool {
	if o != nil && !IsNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given ContainerGroupProbeHttp and assigns it to the Http field.
func (o *ContainerGroupReadinessProbe) SetHttp(v ContainerGroupProbeHttp) {
	o.Http = &v
}

// GetGrpc returns the Grpc field value if set, zero value otherwise.
func (o *ContainerGroupReadinessProbe) GetGrpc() ContainerGroupProbeGrpc {
	if o == nil || IsNil(o.Grpc) {
		var ret ContainerGroupProbeGrpc
		return ret
	}
	return *o.Grpc
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetGrpcOk() (*ContainerGroupProbeGrpc, bool) {
	if o == nil || IsNil(o.Grpc) {
		return nil, false
	}
	return o.Grpc, true
}

// HasGrpc returns a boolean if a field has been set.
func (o *ContainerGroupReadinessProbe) HasGrpc() bool {
	if o != nil && !IsNil(o.Grpc) {
		return true
	}

	return false
}

// SetGrpc gets a reference to the given ContainerGroupProbeGrpc and assigns it to the Grpc field.
func (o *ContainerGroupReadinessProbe) SetGrpc(v ContainerGroupProbeGrpc) {
	o.Grpc = &v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *ContainerGroupReadinessProbe) GetExec() ContainerGroupProbeExec {
	if o == nil || IsNil(o.Exec) {
		var ret ContainerGroupProbeExec
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetExecOk() (*ContainerGroupProbeExec, bool) {
	if o == nil || IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *ContainerGroupReadinessProbe) HasExec() bool {
	if o != nil && !IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given ContainerGroupProbeExec and assigns it to the Exec field.
func (o *ContainerGroupReadinessProbe) SetExec(v ContainerGroupProbeExec) {
	o.Exec = &v
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value
func (o *ContainerGroupReadinessProbe) GetInitialDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InitialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialDelaySeconds, true
}

// SetInitialDelaySeconds sets field value
func (o *ContainerGroupReadinessProbe) SetInitialDelaySeconds(v int32) {
	o.InitialDelaySeconds = v
}

// GetPeriodSeconds returns the PeriodSeconds field value
func (o *ContainerGroupReadinessProbe) GetPeriodSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeriodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodSeconds, true
}

// SetPeriodSeconds sets field value
func (o *ContainerGroupReadinessProbe) SetPeriodSeconds(v int32) {
	o.PeriodSeconds = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value
func (o *ContainerGroupReadinessProbe) GetTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutSeconds, true
}

// SetTimeoutSeconds sets field value
func (o *ContainerGroupReadinessProbe) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = v
}

// GetSuccessThreshold returns the SuccessThreshold field value
func (o *ContainerGroupReadinessProbe) GetSuccessThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SuccessThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessThreshold, true
}

// SetSuccessThreshold sets field value
func (o *ContainerGroupReadinessProbe) SetSuccessThreshold(v int32) {
	o.SuccessThreshold = v
}

// GetFailureThreshold returns the FailureThreshold field value
func (o *ContainerGroupReadinessProbe) GetFailureThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupReadinessProbe) GetFailureThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureThreshold, true
}

// SetFailureThreshold sets field value
func (o *ContainerGroupReadinessProbe) SetFailureThreshold(v int32) {
	o.FailureThreshold = v
}

func (o ContainerGroupReadinessProbe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupReadinessProbe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tcp) {
		toSerialize["tcp"] = o.Tcp
	}
	if !IsNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	if !IsNil(o.Grpc) {
		toSerialize["grpc"] = o.Grpc
	}
	if !IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	toSerialize["initial_delay_seconds"] = o.InitialDelaySeconds
	toSerialize["period_seconds"] = o.PeriodSeconds
	toSerialize["timeout_seconds"] = o.TimeoutSeconds
	toSerialize["success_threshold"] = o.SuccessThreshold
	toSerialize["failure_threshold"] = o.FailureThreshold
	return toSerialize, nil
}

func (o *ContainerGroupReadinessProbe) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"initial_delay_seconds",
		"period_seconds",
		"timeout_seconds",
		"success_threshold",
		"failure_threshold",
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

	varContainerGroupReadinessProbe := _ContainerGroupReadinessProbe{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroupReadinessProbe)

	if err != nil {
		return err
	}

	*o = ContainerGroupReadinessProbe(varContainerGroupReadinessProbe)

	return err
}

type NullableContainerGroupReadinessProbe struct {
	value *ContainerGroupReadinessProbe
	isSet bool
}

func (v NullableContainerGroupReadinessProbe) Get() *ContainerGroupReadinessProbe {
	return v.value
}

func (v *NullableContainerGroupReadinessProbe) Set(val *ContainerGroupReadinessProbe) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupReadinessProbe) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupReadinessProbe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupReadinessProbe(val *ContainerGroupReadinessProbe) *NullableContainerGroupReadinessProbe {
	return &NullableContainerGroupReadinessProbe{value: val, isSet: true}
}

func (v NullableContainerGroupReadinessProbe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupReadinessProbe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


