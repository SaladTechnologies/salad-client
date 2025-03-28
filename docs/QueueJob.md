# QueueJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The job identifier | 
**Input** | **interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata for the job | [optional] 
**Webhook** | Pointer to **string** | The webhook URL to notify when the job completes | [optional] 
**Status** | **string** | The job status | 
**Events** | [**[]QueueJobEvent**](QueueJobEvent.md) | The job events | 
**Output** | Pointer to **interface{}** |  | [optional] 
**CreateTime** | **time.Time** | The job creation time | 
**UpdateTime** | **time.Time** | The job update time | 

## Methods

### NewQueueJob

`func NewQueueJob(id string, input interface{}, status string, events []QueueJobEvent, createTime time.Time, updateTime time.Time, ) *QueueJob`

NewQueueJob instantiates a new QueueJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueJobWithDefaults

`func NewQueueJobWithDefaults() *QueueJob`

NewQueueJobWithDefaults instantiates a new QueueJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueJob) SetId(v string)`

SetId sets Id field to given value.


### GetInput

`func (o *QueueJob) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *QueueJob) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *QueueJob) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *QueueJob) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *QueueJob) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *QueueJob) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueueJob) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueueJob) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueueJob) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhook

`func (o *QueueJob) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *QueueJob) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *QueueJob) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *QueueJob) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetStatus

`func (o *QueueJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueueJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueueJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEvents

`func (o *QueueJob) GetEvents() []QueueJobEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *QueueJob) GetEventsOk() (*[]QueueJobEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *QueueJob) SetEvents(v []QueueJobEvent)`

SetEvents sets Events field to given value.


### GetOutput

`func (o *QueueJob) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *QueueJob) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *QueueJob) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *QueueJob) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *QueueJob) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *QueueJob) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetCreateTime

`func (o *QueueJob) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueueJob) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueueJob) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *QueueJob) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueueJob) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueueJob) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


