# ContainerGroupInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | The unique instance ID | 
**MachineId** | **string** | The machine ID | 
**State** | **string** | The state of the container group instance | 
**UpdateTime** | **time.Time** | The UTC date &amp; time when the workload on this machine transitioned to the current state | 
**Version** | **int32** | The version of the running container group | 
**Ready** | Pointer to **bool** | Specifies whether the container group instance is currently passing its readiness check. If no readiness probe is defined, is true once fully started. | [optional] 
**Started** | Pointer to **bool** | Specifies whether the container group instance passed its startup probe. Is always true when no startup probe is defined. | [optional] 

## Methods

### NewContainerGroupInstance

`func NewContainerGroupInstance(instanceId string, machineId string, state string, updateTime time.Time, version int32, ) *ContainerGroupInstance`

NewContainerGroupInstance instantiates a new ContainerGroupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupInstanceWithDefaults

`func NewContainerGroupInstanceWithDefaults() *ContainerGroupInstance`

NewContainerGroupInstanceWithDefaults instantiates a new ContainerGroupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ContainerGroupInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ContainerGroupInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ContainerGroupInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMachineId

`func (o *ContainerGroupInstance) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *ContainerGroupInstance) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *ContainerGroupInstance) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.


### GetState

`func (o *ContainerGroupInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContainerGroupInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContainerGroupInstance) SetState(v string)`

SetState sets State field to given value.


### GetUpdateTime

`func (o *ContainerGroupInstance) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContainerGroupInstance) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContainerGroupInstance) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetVersion

`func (o *ContainerGroupInstance) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContainerGroupInstance) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContainerGroupInstance) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetReady

`func (o *ContainerGroupInstance) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ContainerGroupInstance) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ContainerGroupInstance) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *ContainerGroupInstance) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetStarted

`func (o *ContainerGroupInstance) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *ContainerGroupInstance) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *ContainerGroupInstance) SetStarted(v bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *ContainerGroupInstance) HasStarted() bool`

HasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


