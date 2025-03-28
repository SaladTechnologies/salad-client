# ContainerGroupInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The container group instance identifier. | 
**MachineId** | **string** | The container group machine identifier. | 
**State** | [**ContainerGroupInstanceState**](ContainerGroupInstanceState.md) |  | 
**UpdateTime** | **time.Time** | The UTC timestamp when the container group instance last changed its state. This helps track the lifecycle and state transitions of the instance. | 
**Version** | **int32** | The version of the container group definition currently running on this instance. Used to track deployment and update progress across the container group fleet. | 
**Ready** | Pointer to **bool** | Indicates whether the container group instance is currently passing its readiness checks and is able to receive traffic or perform its intended function. If no readiness probe is defined, this will be true once the instance is fully started. | [optional] 
**Started** | Pointer to **bool** | Indicates whether the container group instance has successfully completed its startup sequence and passed any configured startup probes. This will always be true when no startup probe is defined for the container group. | [optional] 
**DeletionCost** | Pointer to **int32** | The cost of deleting the container group instance | [optional] [default to 0]

## Methods

### NewContainerGroupInstance

`func NewContainerGroupInstance(id string, machineId string, state ContainerGroupInstanceState, updateTime time.Time, version int32, ) *ContainerGroupInstance`

NewContainerGroupInstance instantiates a new ContainerGroupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupInstanceWithDefaults

`func NewContainerGroupInstanceWithDefaults() *ContainerGroupInstance`

NewContainerGroupInstanceWithDefaults instantiates a new ContainerGroupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerGroupInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerGroupInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerGroupInstance) SetId(v string)`

SetId sets Id field to given value.


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

`func (o *ContainerGroupInstance) GetState() ContainerGroupInstanceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContainerGroupInstance) GetStateOk() (*ContainerGroupInstanceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContainerGroupInstance) SetState(v ContainerGroupInstanceState)`

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

### GetDeletionCost

`func (o *ContainerGroupInstance) GetDeletionCost() int32`

GetDeletionCost returns the DeletionCost field if non-nil, zero value otherwise.

### GetDeletionCostOk

`func (o *ContainerGroupInstance) GetDeletionCostOk() (*int32, bool)`

GetDeletionCostOk returns a tuple with the DeletionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionCost

`func (o *ContainerGroupInstance) SetDeletionCost(v int32)`

SetDeletionCost sets DeletionCost field to given value.

### HasDeletionCost

`func (o *ContainerGroupInstance) HasDeletionCost() bool`

HasDeletionCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


