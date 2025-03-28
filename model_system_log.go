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
	"time"
	"bytes"
	"fmt"
)

// checks if the SystemLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemLog{}

// SystemLog Represents a system log
type SystemLog struct {
	// The name of the event
	EventName string `json:"event_name" validate:"regexp=^.*$"`
	// The UTC date & time when the log item was created
	EventTime time.Time `json:"event_time"`
	// The container group instance identifier.
	InstanceId *string `json:"instance_id,omitempty"`
	// The container group machine identifier.
	MachineId *string `json:"machine_id,omitempty"`
	// The number of CPUs
	ResourceCpu NullableInt32 `json:"resource_cpu"`
	// The GPU class name
	ResourceGpuClass string `json:"resource_gpu_class"`
	// The memory amount in MB
	ResourceMemory NullableInt32 `json:"resource_memory"`
	// The storage amount in bytes
	ResourceStorageAmount NullableInt64 `json:"resource_storage_amount"`
	// The version instance ID
	Version string `json:"version"`
}

type _SystemLog SystemLog

// NewSystemLog instantiates a new SystemLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemLog(eventName string, eventTime time.Time, resourceCpu NullableInt32, resourceGpuClass string, resourceMemory NullableInt32, resourceStorageAmount NullableInt64, version string) *SystemLog {
	this := SystemLog{}
	this.EventName = eventName
	this.EventTime = eventTime
	this.ResourceCpu = resourceCpu
	this.ResourceGpuClass = resourceGpuClass
	this.ResourceMemory = resourceMemory
	this.ResourceStorageAmount = resourceStorageAmount
	this.Version = version
	return &this
}

// NewSystemLogWithDefaults instantiates a new SystemLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemLogWithDefaults() *SystemLog {
	this := SystemLog{}
	return &this
}

// GetEventName returns the EventName field value
func (o *SystemLog) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *SystemLog) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *SystemLog) SetEventName(v string) {
	o.EventName = v
}

// GetEventTime returns the EventTime field value
func (o *SystemLog) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *SystemLog) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *SystemLog) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *SystemLog) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *SystemLog) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *SystemLog) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetMachineId returns the MachineId field value if set, zero value otherwise.
func (o *SystemLog) GetMachineId() string {
	if o == nil || IsNil(o.MachineId) {
		var ret string
		return ret
	}
	return *o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemLog) GetMachineIdOk() (*string, bool) {
	if o == nil || IsNil(o.MachineId) {
		return nil, false
	}
	return o.MachineId, true
}

// HasMachineId returns a boolean if a field has been set.
func (o *SystemLog) HasMachineId() bool {
	if o != nil && !IsNil(o.MachineId) {
		return true
	}

	return false
}

// SetMachineId gets a reference to the given string and assigns it to the MachineId field.
func (o *SystemLog) SetMachineId(v string) {
	o.MachineId = &v
}

// GetResourceCpu returns the ResourceCpu field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SystemLog) GetResourceCpu() int32 {
	if o == nil || o.ResourceCpu.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ResourceCpu.Get()
}

// GetResourceCpuOk returns a tuple with the ResourceCpu field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemLog) GetResourceCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceCpu.Get(), o.ResourceCpu.IsSet()
}

// SetResourceCpu sets field value
func (o *SystemLog) SetResourceCpu(v int32) {
	o.ResourceCpu.Set(&v)
}

// GetResourceGpuClass returns the ResourceGpuClass field value
func (o *SystemLog) GetResourceGpuClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGpuClass
}

// GetResourceGpuClassOk returns a tuple with the ResourceGpuClass field value
// and a boolean to check if the value has been set.
func (o *SystemLog) GetResourceGpuClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGpuClass, true
}

// SetResourceGpuClass sets field value
func (o *SystemLog) SetResourceGpuClass(v string) {
	o.ResourceGpuClass = v
}

// GetResourceMemory returns the ResourceMemory field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SystemLog) GetResourceMemory() int32 {
	if o == nil || o.ResourceMemory.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ResourceMemory.Get()
}

// GetResourceMemoryOk returns a tuple with the ResourceMemory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemLog) GetResourceMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceMemory.Get(), o.ResourceMemory.IsSet()
}

// SetResourceMemory sets field value
func (o *SystemLog) SetResourceMemory(v int32) {
	o.ResourceMemory.Set(&v)
}

// GetResourceStorageAmount returns the ResourceStorageAmount field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *SystemLog) GetResourceStorageAmount() int64 {
	if o == nil || o.ResourceStorageAmount.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ResourceStorageAmount.Get()
}

// GetResourceStorageAmountOk returns a tuple with the ResourceStorageAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemLog) GetResourceStorageAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceStorageAmount.Get(), o.ResourceStorageAmount.IsSet()
}

// SetResourceStorageAmount sets field value
func (o *SystemLog) SetResourceStorageAmount(v int64) {
	o.ResourceStorageAmount.Set(&v)
}

// GetVersion returns the Version field value
func (o *SystemLog) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SystemLog) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SystemLog) SetVersion(v string) {
	o.Version = v
}

func (o SystemLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_name"] = o.EventName
	toSerialize["event_time"] = o.EventTime
	if !IsNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !IsNil(o.MachineId) {
		toSerialize["machine_id"] = o.MachineId
	}
	toSerialize["resource_cpu"] = o.ResourceCpu.Get()
	toSerialize["resource_gpu_class"] = o.ResourceGpuClass
	toSerialize["resource_memory"] = o.ResourceMemory.Get()
	toSerialize["resource_storage_amount"] = o.ResourceStorageAmount.Get()
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *SystemLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_name",
		"event_time",
		"resource_cpu",
		"resource_gpu_class",
		"resource_memory",
		"resource_storage_amount",
		"version",
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

	varSystemLog := _SystemLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemLog)

	if err != nil {
		return err
	}

	*o = SystemLog(varSystemLog)

	return err
}

type NullableSystemLog struct {
	value *SystemLog
	isSet bool
}

func (v NullableSystemLog) Get() *SystemLog {
	return v.value
}

func (v *NullableSystemLog) Set(val *SystemLog) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemLog) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemLog(val *SystemLog) *NullableSystemLog {
	return &NullableSystemLog{value: val, isSet: true}
}

func (v NullableSystemLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


