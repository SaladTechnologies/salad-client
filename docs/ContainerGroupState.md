# ContainerGroupState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ContainerGroupStatus**](ContainerGroupStatus.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartTime** | **time.Time** |  | 
**FinishTime** | **time.Time** |  | 
**InstanceStatusCounts** | [**ContainerGroupInstanceStatusCount**](ContainerGroupInstanceStatusCount.md) |  | 

## Methods

### NewContainerGroupState

`func NewContainerGroupState(status ContainerGroupStatus, startTime time.Time, finishTime time.Time, instanceStatusCounts ContainerGroupInstanceStatusCount, ) *ContainerGroupState`

NewContainerGroupState instantiates a new ContainerGroupState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupStateWithDefaults

`func NewContainerGroupStateWithDefaults() *ContainerGroupState`

NewContainerGroupStateWithDefaults instantiates a new ContainerGroupState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ContainerGroupState) GetStatus() ContainerGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContainerGroupState) GetStatusOk() (*ContainerGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContainerGroupState) SetStatus(v ContainerGroupStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ContainerGroupState) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerGroupState) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerGroupState) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerGroupState) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContainerGroupState) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerGroupState) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartTime

`func (o *ContainerGroupState) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ContainerGroupState) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ContainerGroupState) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetFinishTime

`func (o *ContainerGroupState) GetFinishTime() time.Time`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ContainerGroupState) GetFinishTimeOk() (*time.Time, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ContainerGroupState) SetFinishTime(v time.Time)`

SetFinishTime sets FinishTime field to given value.


### GetInstanceStatusCounts

`func (o *ContainerGroupState) GetInstanceStatusCounts() ContainerGroupInstanceStatusCount`

GetInstanceStatusCounts returns the InstanceStatusCounts field if non-nil, zero value otherwise.

### GetInstanceStatusCountsOk

`func (o *ContainerGroupState) GetInstanceStatusCountsOk() (*ContainerGroupInstanceStatusCount, bool)`

GetInstanceStatusCountsOk returns a tuple with the InstanceStatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStatusCounts

`func (o *ContainerGroupState) SetInstanceStatusCounts(v ContainerGroupInstanceStatusCount)`

SetInstanceStatusCounts sets InstanceStatusCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


