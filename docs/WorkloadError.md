# WorkloadError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**FailedAt** | **time.Time** |  | 
**InstanceId** | **string** |  | 
**MachineId** | **string** |  | 
**AllocatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**Version** | **int32** |  | 

## Methods

### NewWorkloadError

`func NewWorkloadError(detail string, failedAt time.Time, instanceId string, machineId string, allocatedAt time.Time, version int32, ) *WorkloadError`

NewWorkloadError instantiates a new WorkloadError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadErrorWithDefaults

`func NewWorkloadErrorWithDefaults() *WorkloadError`

NewWorkloadErrorWithDefaults instantiates a new WorkloadError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *WorkloadError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *WorkloadError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *WorkloadError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetFailedAt

`func (o *WorkloadError) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *WorkloadError) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *WorkloadError) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.


### GetInstanceId

`func (o *WorkloadError) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *WorkloadError) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *WorkloadError) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMachineId

`func (o *WorkloadError) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *WorkloadError) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *WorkloadError) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.


### GetAllocatedAt

`func (o *WorkloadError) GetAllocatedAt() time.Time`

GetAllocatedAt returns the AllocatedAt field if non-nil, zero value otherwise.

### GetAllocatedAtOk

`func (o *WorkloadError) GetAllocatedAtOk() (*time.Time, bool)`

GetAllocatedAtOk returns a tuple with the AllocatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedAt

`func (o *WorkloadError) SetAllocatedAt(v time.Time)`

SetAllocatedAt sets AllocatedAt field to given value.


### GetStartedAt

`func (o *WorkloadError) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkloadError) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkloadError) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkloadError) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *WorkloadError) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *WorkloadError) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetVersion

`func (o *WorkloadError) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadError) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadError) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


