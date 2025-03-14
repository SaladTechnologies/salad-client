# InferenceEndpointJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The inference endpoint job identifier. | 
**InferenceEndpointName** | **string** | The inference endpoint name. | 
**OrganizationName** | **string** | The organization name. | 
**Input** | **interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** | The job metadata. May be any valid JSON. | [optional] 
**Webhook** | Pointer to **string** | The webhook URL called when the job completes. | [optional] 
**WebhookUrl** | Pointer to **string** | The webhook URL called when the job completes. | [optional] 
**Status** | [**InferenceEndpointJobStatus**](InferenceEndpointJobStatus.md) |  | 
**Events** | [**[]InferenceEndpointJobEvent**](InferenceEndpointJobEvent.md) | The list of events. | 
**Output** | Pointer to **interface{}** |  | [optional] 
**CreateTime** | **time.Time** | The time the job was created. | 
**UpdateTime** | **time.Time** | The time the job was last updated. | 

## Methods

### NewInferenceEndpointJob

`func NewInferenceEndpointJob(id string, inferenceEndpointName string, organizationName string, input interface{}, status InferenceEndpointJobStatus, events []InferenceEndpointJobEvent, createTime time.Time, updateTime time.Time, ) *InferenceEndpointJob`

NewInferenceEndpointJob instantiates a new InferenceEndpointJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceEndpointJobWithDefaults

`func NewInferenceEndpointJobWithDefaults() *InferenceEndpointJob`

NewInferenceEndpointJobWithDefaults instantiates a new InferenceEndpointJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InferenceEndpointJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InferenceEndpointJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InferenceEndpointJob) SetId(v string)`

SetId sets Id field to given value.


### GetInferenceEndpointName

`func (o *InferenceEndpointJob) GetInferenceEndpointName() string`

GetInferenceEndpointName returns the InferenceEndpointName field if non-nil, zero value otherwise.

### GetInferenceEndpointNameOk

`func (o *InferenceEndpointJob) GetInferenceEndpointNameOk() (*string, bool)`

GetInferenceEndpointNameOk returns a tuple with the InferenceEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceEndpointName

`func (o *InferenceEndpointJob) SetInferenceEndpointName(v string)`

SetInferenceEndpointName sets InferenceEndpointName field to given value.


### GetOrganizationName

`func (o *InferenceEndpointJob) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InferenceEndpointJob) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InferenceEndpointJob) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetInput

`func (o *InferenceEndpointJob) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InferenceEndpointJob) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InferenceEndpointJob) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *InferenceEndpointJob) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *InferenceEndpointJob) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *InferenceEndpointJob) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InferenceEndpointJob) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InferenceEndpointJob) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InferenceEndpointJob) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhook

`func (o *InferenceEndpointJob) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *InferenceEndpointJob) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *InferenceEndpointJob) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *InferenceEndpointJob) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *InferenceEndpointJob) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *InferenceEndpointJob) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *InferenceEndpointJob) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *InferenceEndpointJob) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InferenceEndpointJob) GetStatus() InferenceEndpointJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InferenceEndpointJob) GetStatusOk() (*InferenceEndpointJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InferenceEndpointJob) SetStatus(v InferenceEndpointJobStatus)`

SetStatus sets Status field to given value.


### GetEvents

`func (o *InferenceEndpointJob) GetEvents() []InferenceEndpointJobEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InferenceEndpointJob) GetEventsOk() (*[]InferenceEndpointJobEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *InferenceEndpointJob) SetEvents(v []InferenceEndpointJobEvent)`

SetEvents sets Events field to given value.


### GetOutput

`func (o *InferenceEndpointJob) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *InferenceEndpointJob) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *InferenceEndpointJob) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *InferenceEndpointJob) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *InferenceEndpointJob) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *InferenceEndpointJob) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetCreateTime

`func (o *InferenceEndpointJob) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *InferenceEndpointJob) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *InferenceEndpointJob) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *InferenceEndpointJob) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *InferenceEndpointJob) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *InferenceEndpointJob) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


