# ContainerLoggingDatadog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The Datadog intake server host URL where logs will be sent. | 
**ApiKey** | **string** | The Datadog API key used for authentication when sending logs. | 
**Tags** | [**[]ContainerLoggingDatadogTag**](ContainerLoggingDatadogTag.md) | Optional metadata tags to attach to logs for filtering and categorization in Datadog. | 

## Methods

### NewContainerLoggingDatadog

`func NewContainerLoggingDatadog(host string, apiKey string, tags []ContainerLoggingDatadogTag, ) *ContainerLoggingDatadog`

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

`func (o *ContainerLoggingDatadog) GetTags() []ContainerLoggingDatadogTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContainerLoggingDatadog) GetTagsOk() (*[]ContainerLoggingDatadogTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContainerLoggingDatadog) SetTags(v []ContainerLoggingDatadogTag)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ContainerLoggingDatadog) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ContainerLoggingDatadog) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


