# QueuePrototype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The queue name. This must be unique within the project. | 
**DisplayName** | Pointer to **string** | The display name. This may be used as a more human-readable name. | [optional] 
**Description** | Pointer to **string** | The description. This may be used as a space for notes or other information about the queue. | [optional] 

## Methods

### NewQueuePrototype

`func NewQueuePrototype(name string, ) *QueuePrototype`

NewQueuePrototype instantiates a new QueuePrototype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuePrototypeWithDefaults

`func NewQueuePrototypeWithDefaults() *QueuePrototype`

NewQueuePrototypeWithDefaults instantiates a new QueuePrototype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueuePrototype) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueuePrototype) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueuePrototype) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *QueuePrototype) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *QueuePrototype) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *QueuePrototype) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *QueuePrototype) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *QueuePrototype) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueuePrototype) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueuePrototype) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueuePrototype) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


