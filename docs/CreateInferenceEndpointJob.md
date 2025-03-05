# CreateInferenceEndpointJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Webhook** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateInferenceEndpointJob

`func NewCreateInferenceEndpointJob(input interface{}, ) *CreateInferenceEndpointJob`

NewCreateInferenceEndpointJob instantiates a new CreateInferenceEndpointJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInferenceEndpointJobWithDefaults

`func NewCreateInferenceEndpointJobWithDefaults() *CreateInferenceEndpointJob`

NewCreateInferenceEndpointJobWithDefaults instantiates a new CreateInferenceEndpointJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *CreateInferenceEndpointJob) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CreateInferenceEndpointJob) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CreateInferenceEndpointJob) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *CreateInferenceEndpointJob) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *CreateInferenceEndpointJob) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *CreateInferenceEndpointJob) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateInferenceEndpointJob) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateInferenceEndpointJob) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateInferenceEndpointJob) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhook

`func (o *CreateInferenceEndpointJob) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CreateInferenceEndpointJob) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CreateInferenceEndpointJob) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CreateInferenceEndpointJob) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateInferenceEndpointJob) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateInferenceEndpointJob) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateInferenceEndpointJob) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateInferenceEndpointJob) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


