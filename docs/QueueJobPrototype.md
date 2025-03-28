# QueueJobPrototype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **interface{}** |  | 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata for the job | [optional] 
**Webhook** | Pointer to **string** | The webhook to call when the job completes | [optional] 

## Methods

### NewQueueJobPrototype

`func NewQueueJobPrototype(input interface{}, ) *QueueJobPrototype`

NewQueueJobPrototype instantiates a new QueueJobPrototype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueJobPrototypeWithDefaults

`func NewQueueJobPrototypeWithDefaults() *QueueJobPrototype`

NewQueueJobPrototypeWithDefaults instantiates a new QueueJobPrototype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *QueueJobPrototype) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *QueueJobPrototype) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *QueueJobPrototype) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *QueueJobPrototype) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *QueueJobPrototype) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetMetadata

`func (o *QueueJobPrototype) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueueJobPrototype) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueueJobPrototype) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueueJobPrototype) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhook

`func (o *QueueJobPrototype) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *QueueJobPrototype) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *QueueJobPrototype) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *QueueJobPrototype) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


