# SystemLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of the event | 
**EventTime** | **time.Time** | The UTC date &amp; time when the log item was created | 
**InstanceId** | Pointer to **string** | The container group instance identifier. | [optional] 
**MachineId** | Pointer to **string** | The container group machine identifier. | [optional] 
**ResourceCpu** | **NullableInt32** | The number of CPUs | 
**ResourceGpuClass** | **string** | The GPU class name | 
**ResourceMemory** | **NullableInt32** | The memory amount in MB | 
**ResourceStorageAmount** | **NullableInt64** | The storage amount in bytes | 
**Version** | **string** | The version instance ID | 

## Methods

### NewSystemLog

`func NewSystemLog(eventName string, eventTime time.Time, resourceCpu NullableInt32, resourceGpuClass string, resourceMemory NullableInt32, resourceStorageAmount NullableInt64, version string, ) *SystemLog`

NewSystemLog instantiates a new SystemLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLogWithDefaults

`func NewSystemLogWithDefaults() *SystemLog`

NewSystemLogWithDefaults instantiates a new SystemLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *SystemLog) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *SystemLog) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *SystemLog) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *SystemLog) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SystemLog) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SystemLog) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetInstanceId

`func (o *SystemLog) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *SystemLog) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *SystemLog) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *SystemLog) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetMachineId

`func (o *SystemLog) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *SystemLog) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *SystemLog) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *SystemLog) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetResourceCpu

`func (o *SystemLog) GetResourceCpu() int32`

GetResourceCpu returns the ResourceCpu field if non-nil, zero value otherwise.

### GetResourceCpuOk

`func (o *SystemLog) GetResourceCpuOk() (*int32, bool)`

GetResourceCpuOk returns a tuple with the ResourceCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCpu

`func (o *SystemLog) SetResourceCpu(v int32)`

SetResourceCpu sets ResourceCpu field to given value.


### SetResourceCpuNil

`func (o *SystemLog) SetResourceCpuNil(b bool)`

 SetResourceCpuNil sets the value for ResourceCpu to be an explicit nil

### UnsetResourceCpu
`func (o *SystemLog) UnsetResourceCpu()`

UnsetResourceCpu ensures that no value is present for ResourceCpu, not even an explicit nil
### GetResourceGpuClass

`func (o *SystemLog) GetResourceGpuClass() string`

GetResourceGpuClass returns the ResourceGpuClass field if non-nil, zero value otherwise.

### GetResourceGpuClassOk

`func (o *SystemLog) GetResourceGpuClassOk() (*string, bool)`

GetResourceGpuClassOk returns a tuple with the ResourceGpuClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGpuClass

`func (o *SystemLog) SetResourceGpuClass(v string)`

SetResourceGpuClass sets ResourceGpuClass field to given value.


### GetResourceMemory

`func (o *SystemLog) GetResourceMemory() int32`

GetResourceMemory returns the ResourceMemory field if non-nil, zero value otherwise.

### GetResourceMemoryOk

`func (o *SystemLog) GetResourceMemoryOk() (*int32, bool)`

GetResourceMemoryOk returns a tuple with the ResourceMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMemory

`func (o *SystemLog) SetResourceMemory(v int32)`

SetResourceMemory sets ResourceMemory field to given value.


### SetResourceMemoryNil

`func (o *SystemLog) SetResourceMemoryNil(b bool)`

 SetResourceMemoryNil sets the value for ResourceMemory to be an explicit nil

### UnsetResourceMemory
`func (o *SystemLog) UnsetResourceMemory()`

UnsetResourceMemory ensures that no value is present for ResourceMemory, not even an explicit nil
### GetResourceStorageAmount

`func (o *SystemLog) GetResourceStorageAmount() int64`

GetResourceStorageAmount returns the ResourceStorageAmount field if non-nil, zero value otherwise.

### GetResourceStorageAmountOk

`func (o *SystemLog) GetResourceStorageAmountOk() (*int64, bool)`

GetResourceStorageAmountOk returns a tuple with the ResourceStorageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStorageAmount

`func (o *SystemLog) SetResourceStorageAmount(v int64)`

SetResourceStorageAmount sets ResourceStorageAmount field to given value.


### SetResourceStorageAmountNil

`func (o *SystemLog) SetResourceStorageAmountNil(b bool)`

 SetResourceStorageAmountNil sets the value for ResourceStorageAmount to be an explicit nil

### UnsetResourceStorageAmount
`func (o *SystemLog) UnsetResourceStorageAmount()`

UnsetResourceStorageAmount ensures that no value is present for ResourceStorageAmount, not even an explicit nil
### GetVersion

`func (o *SystemLog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SystemLog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SystemLog) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


