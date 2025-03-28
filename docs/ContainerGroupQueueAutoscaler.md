# ContainerGroupQueueAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredQueueLength** | **int32** | The target number of items in the queue that the autoscaler attempts to maintain by scaling the containers up or down | 
**MaxReplicas** | **int32** | The maximum number of instances the container can scale up to | 
**MaxDownscalePerMinute** | Pointer to **int32** | The maximum number of instances that can be removed per minute to prevent rapid downscaling | [optional] 
**MaxUpscalePerMinute** | Pointer to **int32** | The maximum number of instances that can be added per minute to prevent rapid upscaling | [optional] 
**MinReplicas** | **int32** | The minimum number of instances the container can scale down to, ensuring baseline availability | 
**PollingPeriod** | Pointer to **int32** | The period (in seconds) in which the autoscaler checks the queue length and applies the scaling formula | [optional] 

## Methods

### NewContainerGroupQueueAutoscaler

`func NewContainerGroupQueueAutoscaler(desiredQueueLength int32, maxReplicas int32, minReplicas int32, ) *ContainerGroupQueueAutoscaler`

NewContainerGroupQueueAutoscaler instantiates a new ContainerGroupQueueAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupQueueAutoscalerWithDefaults

`func NewContainerGroupQueueAutoscalerWithDefaults() *ContainerGroupQueueAutoscaler`

NewContainerGroupQueueAutoscalerWithDefaults instantiates a new ContainerGroupQueueAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredQueueLength

`func (o *ContainerGroupQueueAutoscaler) GetDesiredQueueLength() int32`

GetDesiredQueueLength returns the DesiredQueueLength field if non-nil, zero value otherwise.

### GetDesiredQueueLengthOk

`func (o *ContainerGroupQueueAutoscaler) GetDesiredQueueLengthOk() (*int32, bool)`

GetDesiredQueueLengthOk returns a tuple with the DesiredQueueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQueueLength

`func (o *ContainerGroupQueueAutoscaler) SetDesiredQueueLength(v int32)`

SetDesiredQueueLength sets DesiredQueueLength field to given value.


### GetMaxReplicas

`func (o *ContainerGroupQueueAutoscaler) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *ContainerGroupQueueAutoscaler) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *ContainerGroupQueueAutoscaler) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMaxDownscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) GetMaxDownscalePerMinute() int32`

GetMaxDownscalePerMinute returns the MaxDownscalePerMinute field if non-nil, zero value otherwise.

### GetMaxDownscalePerMinuteOk

`func (o *ContainerGroupQueueAutoscaler) GetMaxDownscalePerMinuteOk() (*int32, bool)`

GetMaxDownscalePerMinuteOk returns a tuple with the MaxDownscalePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDownscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) SetMaxDownscalePerMinute(v int32)`

SetMaxDownscalePerMinute sets MaxDownscalePerMinute field to given value.

### HasMaxDownscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) HasMaxDownscalePerMinute() bool`

HasMaxDownscalePerMinute returns a boolean if a field has been set.

### GetMaxUpscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) GetMaxUpscalePerMinute() int32`

GetMaxUpscalePerMinute returns the MaxUpscalePerMinute field if non-nil, zero value otherwise.

### GetMaxUpscalePerMinuteOk

`func (o *ContainerGroupQueueAutoscaler) GetMaxUpscalePerMinuteOk() (*int32, bool)`

GetMaxUpscalePerMinuteOk returns a tuple with the MaxUpscalePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) SetMaxUpscalePerMinute(v int32)`

SetMaxUpscalePerMinute sets MaxUpscalePerMinute field to given value.

### HasMaxUpscalePerMinute

`func (o *ContainerGroupQueueAutoscaler) HasMaxUpscalePerMinute() bool`

HasMaxUpscalePerMinute returns a boolean if a field has been set.

### GetMinReplicas

`func (o *ContainerGroupQueueAutoscaler) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *ContainerGroupQueueAutoscaler) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *ContainerGroupQueueAutoscaler) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.


### GetPollingPeriod

`func (o *ContainerGroupQueueAutoscaler) GetPollingPeriod() int32`

GetPollingPeriod returns the PollingPeriod field if non-nil, zero value otherwise.

### GetPollingPeriodOk

`func (o *ContainerGroupQueueAutoscaler) GetPollingPeriodOk() (*int32, bool)`

GetPollingPeriodOk returns a tuple with the PollingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingPeriod

`func (o *ContainerGroupQueueAutoscaler) SetPollingPeriod(v int32)`

SetPollingPeriod sets PollingPeriod field to given value.

### HasPollingPeriod

`func (o *ContainerGroupQueueAutoscaler) HasPollingPeriod() bool`

HasPollingPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


