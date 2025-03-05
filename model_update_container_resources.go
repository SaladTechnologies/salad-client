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
)

// checks if the UpdateContainerResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContainerResources{}

// UpdateContainerResources struct for UpdateContainerResources
type UpdateContainerResources struct {
	Cpu *int32 `json:"cpu,omitempty"`
	Memory *int32 `json:"memory,omitempty"`
	GpuClasses []string `json:"gpu_classes,omitempty"`
	StorageAmount *int64 `json:"storage_amount,omitempty"`
}

// NewUpdateContainerResources instantiates a new UpdateContainerResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContainerResources() *UpdateContainerResources {
	this := UpdateContainerResources{}
	return &this
}

// NewUpdateContainerResourcesWithDefaults instantiates a new UpdateContainerResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContainerResourcesWithDefaults() *UpdateContainerResources {
	this := UpdateContainerResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *UpdateContainerResources) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerResources) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *UpdateContainerResources) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *UpdateContainerResources) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *UpdateContainerResources) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerResources) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *UpdateContainerResources) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *UpdateContainerResources) SetMemory(v int32) {
	o.Memory = &v
}

// GetGpuClasses returns the GpuClasses field value if set, zero value otherwise.
func (o *UpdateContainerResources) GetGpuClasses() []string {
	if o == nil || IsNil(o.GpuClasses) {
		var ret []string
		return ret
	}
	return o.GpuClasses
}

// GetGpuClassesOk returns a tuple with the GpuClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerResources) GetGpuClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.GpuClasses) {
		return nil, false
	}
	return o.GpuClasses, true
}

// HasGpuClasses returns a boolean if a field has been set.
func (o *UpdateContainerResources) HasGpuClasses() bool {
	if o != nil && !IsNil(o.GpuClasses) {
		return true
	}

	return false
}

// SetGpuClasses gets a reference to the given []string and assigns it to the GpuClasses field.
func (o *UpdateContainerResources) SetGpuClasses(v []string) {
	o.GpuClasses = v
}

// GetStorageAmount returns the StorageAmount field value if set, zero value otherwise.
func (o *UpdateContainerResources) GetStorageAmount() int64 {
	if o == nil || IsNil(o.StorageAmount) {
		var ret int64
		return ret
	}
	return *o.StorageAmount
}

// GetStorageAmountOk returns a tuple with the StorageAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerResources) GetStorageAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageAmount) {
		return nil, false
	}
	return o.StorageAmount, true
}

// HasStorageAmount returns a boolean if a field has been set.
func (o *UpdateContainerResources) HasStorageAmount() bool {
	if o != nil && !IsNil(o.StorageAmount) {
		return true
	}

	return false
}

// SetStorageAmount gets a reference to the given int64 and assigns it to the StorageAmount field.
func (o *UpdateContainerResources) SetStorageAmount(v int64) {
	o.StorageAmount = &v
}

func (o UpdateContainerResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContainerResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.GpuClasses) {
		toSerialize["gpu_classes"] = o.GpuClasses
	}
	if !IsNil(o.StorageAmount) {
		toSerialize["storage_amount"] = o.StorageAmount
	}
	return toSerialize, nil
}

type NullableUpdateContainerResources struct {
	value *UpdateContainerResources
	isSet bool
}

func (v NullableUpdateContainerResources) Get() *UpdateContainerResources {
	return v.value
}

func (v *NullableUpdateContainerResources) Set(val *UpdateContainerResources) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContainerResources) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContainerResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContainerResources(val *UpdateContainerResources) *NullableUpdateContainerResources {
	return &NullableUpdateContainerResources{value: val, isSet: true}
}

func (v NullableUpdateContainerResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContainerResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


