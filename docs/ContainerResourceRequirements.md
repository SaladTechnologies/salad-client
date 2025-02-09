# ContainerResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** |  | 
**Memory** | **int32** |  | 
**GpuClasses** | Pointer to **[]string** |  | [optional] 
**StorageAmount** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewContainerResourceRequirements

`func NewContainerResourceRequirements(cpu int32, memory int32, ) *ContainerResourceRequirements`

NewContainerResourceRequirements instantiates a new ContainerResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerResourceRequirementsWithDefaults

`func NewContainerResourceRequirementsWithDefaults() *ContainerResourceRequirements`

NewContainerResourceRequirementsWithDefaults instantiates a new ContainerResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ContainerResourceRequirements) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ContainerResourceRequirements) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ContainerResourceRequirements) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ContainerResourceRequirements) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ContainerResourceRequirements) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ContainerResourceRequirements) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetGpuClasses

`func (o *ContainerResourceRequirements) GetGpuClasses() []string`

GetGpuClasses returns the GpuClasses field if non-nil, zero value otherwise.

### GetGpuClassesOk

`func (o *ContainerResourceRequirements) GetGpuClassesOk() (*[]string, bool)`

GetGpuClassesOk returns a tuple with the GpuClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClasses

`func (o *ContainerResourceRequirements) SetGpuClasses(v []string)`

SetGpuClasses sets GpuClasses field to given value.

### HasGpuClasses

`func (o *ContainerResourceRequirements) HasGpuClasses() bool`

HasGpuClasses returns a boolean if a field has been set.

### SetGpuClassesNil

`func (o *ContainerResourceRequirements) SetGpuClassesNil(b bool)`

 SetGpuClassesNil sets the value for GpuClasses to be an explicit nil

### UnsetGpuClasses
`func (o *ContainerResourceRequirements) UnsetGpuClasses()`

UnsetGpuClasses ensures that no value is present for GpuClasses, not even an explicit nil
### GetStorageAmount

`func (o *ContainerResourceRequirements) GetStorageAmount() int64`

GetStorageAmount returns the StorageAmount field if non-nil, zero value otherwise.

### GetStorageAmountOk

`func (o *ContainerResourceRequirements) GetStorageAmountOk() (*int64, bool)`

GetStorageAmountOk returns a tuple with the StorageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAmount

`func (o *ContainerResourceRequirements) SetStorageAmount(v int64)`

SetStorageAmount sets StorageAmount field to given value.

### HasStorageAmount

`func (o *ContainerResourceRequirements) HasStorageAmount() bool`

HasStorageAmount returns a boolean if a field has been set.

### SetStorageAmountNil

`func (o *ContainerResourceRequirements) SetStorageAmountNil(b bool)`

 SetStorageAmountNil sets the value for StorageAmount to be an explicit nil

### UnsetStorageAmount
`func (o *ContainerResourceRequirements) UnsetStorageAmount()`

UnsetStorageAmount ensures that no value is present for StorageAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


