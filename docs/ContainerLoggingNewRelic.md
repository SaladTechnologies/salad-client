# ContainerLoggingNewRelic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The New Relic endpoint host for log ingestion (e.g., log-api.newrelic.com). | 
**IngestionKey** | **string** | The New Relic license or ingestion key used for authentication and data routing. | 

## Methods

### NewContainerLoggingNewRelic

`func NewContainerLoggingNewRelic(host string, ingestionKey string, ) *ContainerLoggingNewRelic`

NewContainerLoggingNewRelic instantiates a new ContainerLoggingNewRelic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingNewRelicWithDefaults

`func NewContainerLoggingNewRelicWithDefaults() *ContainerLoggingNewRelic`

NewContainerLoggingNewRelicWithDefaults instantiates a new ContainerLoggingNewRelic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContainerLoggingNewRelic) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContainerLoggingNewRelic) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContainerLoggingNewRelic) SetHost(v string)`

SetHost sets Host field to given value.


### GetIngestionKey

`func (o *ContainerLoggingNewRelic) GetIngestionKey() string`

GetIngestionKey returns the IngestionKey field if non-nil, zero value otherwise.

### GetIngestionKeyOk

`func (o *ContainerLoggingNewRelic) GetIngestionKeyOk() (*string, bool)`

GetIngestionKeyOk returns a tuple with the IngestionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionKey

`func (o *ContainerLoggingNewRelic) SetIngestionKey(v string)`

SetIngestionKey sets IngestionKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


