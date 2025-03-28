# QueueJobEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action that was taken on the queue job | 
**Time** | **time.Time** | The time the action was taken on the queue job | 

## Methods

### NewQueueJobEvent

`func NewQueueJobEvent(action string, time time.Time, ) *QueueJobEvent`

NewQueueJobEvent instantiates a new QueueJobEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueJobEventWithDefaults

`func NewQueueJobEventWithDefaults() *QueueJobEvent`

NewQueueJobEventWithDefaults instantiates a new QueueJobEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *QueueJobEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *QueueJobEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *QueueJobEvent) SetAction(v string)`

SetAction sets Action field to given value.


### GetTime

`func (o *QueueJobEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueueJobEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueueJobEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


