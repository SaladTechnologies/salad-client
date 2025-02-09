# UpdateContainerResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**GpuClasses** | Pointer to **[]string** |  | [optional] 
**StorageAmount** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewUpdateContainerResources

`func NewUpdateContainerResources() *UpdateContainerResources`

NewUpdateContainerResources instantiates a new UpdateContainerResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerResourcesWithDefaults

`func NewUpdateContainerResourcesWithDefaults() *UpdateContainerResources`

NewUpdateContainerResourcesWithDefaults instantiates a new UpdateContainerResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *UpdateContainerResources) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UpdateContainerResources) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UpdateContainerResources) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UpdateContainerResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *UpdateContainerResources) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *UpdateContainerResources) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *UpdateContainerResources) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateContainerResources) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateContainerResources) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateContainerResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *UpdateContainerResources) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *UpdateContainerResources) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetGpuClasses

`func (o *UpdateContainerResources) GetGpuClasses() []string`

GetGpuClasses returns the GpuClasses field if non-nil, zero value otherwise.

### GetGpuClassesOk

`func (o *UpdateContainerResources) GetGpuClassesOk() (*[]string, bool)`

GetGpuClassesOk returns a tuple with the GpuClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClasses

`func (o *UpdateContainerResources) SetGpuClasses(v []string)`

SetGpuClasses sets GpuClasses field to given value.

### HasGpuClasses

`func (o *UpdateContainerResources) HasGpuClasses() bool`

HasGpuClasses returns a boolean if a field has been set.

### SetGpuClassesNil

`func (o *UpdateContainerResources) SetGpuClassesNil(b bool)`

 SetGpuClassesNil sets the value for GpuClasses to be an explicit nil

### UnsetGpuClasses
`func (o *UpdateContainerResources) UnsetGpuClasses()`

UnsetGpuClasses ensures that no value is present for GpuClasses, not even an explicit nil
### GetStorageAmount

`func (o *UpdateContainerResources) GetStorageAmount() int64`

GetStorageAmount returns the StorageAmount field if non-nil, zero value otherwise.

### GetStorageAmountOk

`func (o *UpdateContainerResources) GetStorageAmountOk() (*int64, bool)`

GetStorageAmountOk returns a tuple with the StorageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAmount

`func (o *UpdateContainerResources) SetStorageAmount(v int64)`

SetStorageAmount sets StorageAmount field to given value.

### HasStorageAmount

`func (o *UpdateContainerResources) HasStorageAmount() bool`

HasStorageAmount returns a boolean if a field has been set.

### SetStorageAmountNil

`func (o *UpdateContainerResources) SetStorageAmountNil(b bool)`

 SetStorageAmountNil sets the value for StorageAmount to be an explicit nil

### UnsetStorageAmount
`func (o *UpdateContainerResources) UnsetStorageAmount()`

UnsetStorageAmount ensures that no value is present for StorageAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


