# QueueAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReplicas** | **int32** | The minimum number of instances the container can scale down to | 
**MaxReplicas** | **int32** | The maximum number of instances the container can scale up to | 
**DesiredQueueLength** | **int32** |  | 
**PollingPeriod** | Pointer to **int32** | The period (in seconds) in which the queue checks the formula | [optional] 
**MaxUpscalePerMinute** | Pointer to **int32** | The maximum number of instances that can be added per minute | [optional] 
**MaxDownscalePerMinute** | Pointer to **int32** | The maximum number of instances that can be removed per minute | [optional] 

## Methods

### NewQueueAutoscaler

`func NewQueueAutoscaler(minReplicas int32, maxReplicas int32, desiredQueueLength int32, ) *QueueAutoscaler`

NewQueueAutoscaler instantiates a new QueueAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueAutoscalerWithDefaults

`func NewQueueAutoscalerWithDefaults() *QueueAutoscaler`

NewQueueAutoscalerWithDefaults instantiates a new QueueAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReplicas

`func (o *QueueAutoscaler) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *QueueAutoscaler) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *QueueAutoscaler) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.


### GetMaxReplicas

`func (o *QueueAutoscaler) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *QueueAutoscaler) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *QueueAutoscaler) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetDesiredQueueLength

`func (o *QueueAutoscaler) GetDesiredQueueLength() int32`

GetDesiredQueueLength returns the DesiredQueueLength field if non-nil, zero value otherwise.

### GetDesiredQueueLengthOk

`func (o *QueueAutoscaler) GetDesiredQueueLengthOk() (*int32, bool)`

GetDesiredQueueLengthOk returns a tuple with the DesiredQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQueueLength

`func (o *QueueAutoscaler) SetDesiredQueueLength(v int32)`

SetDesiredQueueLength sets DesiredQueueLength field to given value.


### GetPollingPeriod

`func (o *QueueAutoscaler) GetPollingPeriod() int32`

GetPollingPeriod returns the PollingPeriod field if non-nil, zero value otherwise.

### GetPollingPeriodOk

`func (o *QueueAutoscaler) GetPollingPeriodOk() (*int32, bool)`

GetPollingPeriodOk returns a tuple with the PollingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingPeriod

`func (o *QueueAutoscaler) SetPollingPeriod(v int32)`

SetPollingPeriod sets PollingPeriod field to given value.

### HasPollingPeriod

`func (o *QueueAutoscaler) HasPollingPeriod() bool`

HasPollingPeriod returns a boolean if a field has been set.

### GetMaxUpscalePerMinute

`func (o *QueueAutoscaler) GetMaxUpscalePerMinute() int32`

GetMaxUpscalePerMinute returns the MaxUpscalePerMinute field if non-nil, zero value otherwise.

### GetMaxUpscalePerMinuteOk

`func (o *QueueAutoscaler) GetMaxUpscalePerMinuteOk() (*int32, bool)`

GetMaxUpscalePerMinuteOk returns a tuple with the MaxUpscalePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpscalePerMinute

`func (o *QueueAutoscaler) SetMaxUpscalePerMinute(v int32)`

SetMaxUpscalePerMinute sets MaxUpscalePerMinute field to given value.

### HasMaxUpscalePerMinute

`func (o *QueueAutoscaler) HasMaxUpscalePerMinute() bool`

HasMaxUpscalePerMinute returns a boolean if a field has been set.

### GetMaxDownscalePerMinute

`func (o *QueueAutoscaler) GetMaxDownscalePerMinute() int32`

GetMaxDownscalePerMinute returns the MaxDownscalePerMinute field if non-nil, zero value otherwise.

### GetMaxDownscalePerMinuteOk

`func (o *QueueAutoscaler) GetMaxDownscalePerMinuteOk() (*int32, bool)`

GetMaxDownscalePerMinuteOk returns a tuple with the MaxDownscalePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDownscalePerMinute

`func (o *QueueAutoscaler) SetMaxDownscalePerMinute(v int32)`

SetMaxDownscalePerMinute sets MaxDownscalePerMinute field to given value.

### HasMaxDownscalePerMinute

`func (o *QueueAutoscaler) HasMaxDownscalePerMinute() bool`

HasMaxDownscalePerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


