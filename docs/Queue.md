# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The queue identifier. This is automatically generated and assigned when the queue is created. | 
**Name** | **string** | The queue name. This must be unique within the project. | 
**DisplayName** | **string** | The display name. This may be used as a more human-readable name. | 
**Description** | Pointer to **string** | The description. This may be used as a space for notes or other information about the queue. | [optional] 
**ContainerGroups** | [**[]ContainerGroup**](ContainerGroup.md) | The container groups that are part of this queue. Each container group represents a scalable set of identical containers running as a distributed service. | 
**CreateTime** | **time.Time** | The date and time the queue was created. | 
**UpdateTime** | **time.Time** | The date and time the queue was last updated. | 

## Methods

### NewQueue

`func NewQueue(id string, name string, displayName string, containerGroups []ContainerGroup, createTime time.Time, updateTime time.Time, ) *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Queue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Queue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Queue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Queue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Queue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Queue) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *Queue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Queue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Queue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *Queue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Queue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Queue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Queue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContainerGroups

`func (o *Queue) GetContainerGroups() []ContainerGroup`

GetContainerGroups returns the ContainerGroups field if non-nil, zero value otherwise.

### GetContainerGroupsOk

`func (o *Queue) GetContainerGroupsOk() (*[]ContainerGroup, bool)`

GetContainerGroupsOk returns a tuple with the ContainerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerGroups

`func (o *Queue) SetContainerGroups(v []ContainerGroup)`

SetContainerGroups sets ContainerGroups field to given value.


### GetCreateTime

`func (o *Queue) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Queue) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Queue) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *Queue) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Queue) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Queue) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


