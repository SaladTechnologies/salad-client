# ContainerLoggingDatadog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**ApiKey** | **string** |  | 
**Tags** | Pointer to [**[]ContainerLoggingDatadogTagsInner**](ContainerLoggingDatadogTagsInner.md) |  | [optional] 

## Methods

### NewContainerLoggingDatadog

`func NewContainerLoggingDatadog(host string, apiKey string, ) *ContainerLoggingDatadog`

NewContainerLoggingDatadog instantiates a new ContainerLoggingDatadog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingDatadogWithDefaults

`func NewContainerLoggingDatadogWithDefaults() *ContainerLoggingDatadog`

NewContainerLoggingDatadogWithDefaults instantiates a new ContainerLoggingDatadog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContainerLoggingDatadog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContainerLoggingDatadog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContainerLoggingDatadog) SetHost(v string)`

SetHost sets Host field to given value.


### GetApiKey

`func (o *ContainerLoggingDatadog) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ContainerLoggingDatadog) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ContainerLoggingDatadog) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetTags

`func (o *ContainerLoggingDatadog) GetTags() []ContainerLoggingDatadogTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContainerLoggingDatadog) GetTagsOk() (*[]ContainerLoggingDatadogTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContainerLoggingDatadog) SetTags(v []ContainerLoggingDatadogTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContainerLoggingDatadog) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


