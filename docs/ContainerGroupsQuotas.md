# ContainerGroupsQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerReplicasQuota** | **int32** | The maximum number of replicas that can be created for a container group | 
**ContainerReplicasUsed** | **int32** | The number of replicas that are currently in use | 
**MaxContainerGroupReallocationsPerMinute** | Pointer to **int32** | The maximum number of container group reallocations per minute | [optional] [default to 10]
**MaxContainerGroupRecreatesPerMinute** | Pointer to **int32** | The maximum number of container group recreates per minute | [optional] [default to 10]
**MaxContainerGroupRestartsPerMinute** | Pointer to **int32** | The maximum number of container group restarts per minute | [optional] [default to 10]

## Methods

### NewContainerGroupsQuotas

`func NewContainerGroupsQuotas(containerReplicasQuota int32, containerReplicasUsed int32, ) *ContainerGroupsQuotas`

NewContainerGroupsQuotas instantiates a new ContainerGroupsQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupsQuotasWithDefaults

`func NewContainerGroupsQuotasWithDefaults() *ContainerGroupsQuotas`

NewContainerGroupsQuotasWithDefaults instantiates a new ContainerGroupsQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerReplicasQuota

`func (o *ContainerGroupsQuotas) GetContainerReplicasQuota() int32`

GetContainerReplicasQuota returns the ContainerReplicasQuota field if non-nil, zero value otherwise.

### GetContainerReplicasQuotaOk

`func (o *ContainerGroupsQuotas) GetContainerReplicasQuotaOk() (*int32, bool)`

GetContainerReplicasQuotaOk returns a tuple with the ContainerReplicasQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerReplicasQuota

`func (o *ContainerGroupsQuotas) SetContainerReplicasQuota(v int32)`

SetContainerReplicasQuota sets ContainerReplicasQuota field to given value.


### GetContainerReplicasUsed

`func (o *ContainerGroupsQuotas) GetContainerReplicasUsed() int32`

GetContainerReplicasUsed returns the ContainerReplicasUsed field if non-nil, zero value otherwise.

### GetContainerReplicasUsedOk

`func (o *ContainerGroupsQuotas) GetContainerReplicasUsedOk() (*int32, bool)`

GetContainerReplicasUsedOk returns a tuple with the ContainerReplicasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerReplicasUsed

`func (o *ContainerGroupsQuotas) SetContainerReplicasUsed(v int32)`

SetContainerReplicasUsed sets ContainerReplicasUsed field to given value.


### GetMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinute() int32`

GetMaxContainerGroupReallocationsPerMinute returns the MaxContainerGroupReallocationsPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupReallocationsPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupReallocationsPerMinuteOk returns a tuple with the MaxContainerGroupReallocationsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupReallocationsPerMinute(v int32)`

SetMaxContainerGroupReallocationsPerMinute sets MaxContainerGroupReallocationsPerMinute field to given value.

### HasMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupReallocationsPerMinute() bool`

HasMaxContainerGroupReallocationsPerMinute returns a boolean if a field has been set.

### GetMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinute() int32`

GetMaxContainerGroupRecreatesPerMinute returns the MaxContainerGroupRecreatesPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupRecreatesPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupRecreatesPerMinuteOk returns a tuple with the MaxContainerGroupRecreatesPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupRecreatesPerMinute(v int32)`

SetMaxContainerGroupRecreatesPerMinute sets MaxContainerGroupRecreatesPerMinute field to given value.

### HasMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupRecreatesPerMinute() bool`

HasMaxContainerGroupRecreatesPerMinute returns a boolean if a field has been set.

### GetMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinute() int32`

GetMaxContainerGroupRestartsPerMinute returns the MaxContainerGroupRestartsPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupRestartsPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupRestartsPerMinuteOk returns a tuple with the MaxContainerGroupRestartsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupRestartsPerMinute(v int32)`

SetMaxContainerGroupRestartsPerMinute sets MaxContainerGroupRestartsPerMinute field to given value.

### HasMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupRestartsPerMinute() bool`

HasMaxContainerGroupRestartsPerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


