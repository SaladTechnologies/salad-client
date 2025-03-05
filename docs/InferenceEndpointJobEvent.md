# InferenceEndpointJobEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**InferenceEndpointJobEventAction**](InferenceEndpointJobEventAction.md) |  | 
**Time** | **time.Time** |  | 

## Methods

### NewInferenceEndpointJobEvent

`func NewInferenceEndpointJobEvent(action InferenceEndpointJobEventAction, time time.Time, ) *InferenceEndpointJobEvent`

NewInferenceEndpointJobEvent instantiates a new InferenceEndpointJobEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceEndpointJobEventWithDefaults

`func NewInferenceEndpointJobEventWithDefaults() *InferenceEndpointJobEvent`

NewInferenceEndpointJobEventWithDefaults instantiates a new InferenceEndpointJobEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InferenceEndpointJobEvent) GetAction() InferenceEndpointJobEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InferenceEndpointJobEvent) GetActionOk() (*InferenceEndpointJobEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InferenceEndpointJobEvent) SetAction(v InferenceEndpointJobEventAction)`

SetAction sets Action field to given value.


### GetTime

`func (o *InferenceEndpointJobEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InferenceEndpointJobEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InferenceEndpointJobEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


