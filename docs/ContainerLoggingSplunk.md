# ContainerLoggingSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The URL of the Splunk HTTP Event Collector (HEC) endpoint. | 
**Token** | **string** | The authentication token required to send data to the Splunk HEC endpoint. | 

## Methods

### NewContainerLoggingSplunk

`func NewContainerLoggingSplunk(host string, token string, ) *ContainerLoggingSplunk`

NewContainerLoggingSplunk instantiates a new ContainerLoggingSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingSplunkWithDefaults

`func NewContainerLoggingSplunkWithDefaults() *ContainerLoggingSplunk`

NewContainerLoggingSplunkWithDefaults instantiates a new ContainerLoggingSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContainerLoggingSplunk) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContainerLoggingSplunk) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContainerLoggingSplunk) SetHost(v string)`

SetHost sets Host field to given value.


### GetToken

`func (o *ContainerLoggingSplunk) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ContainerLoggingSplunk) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ContainerLoggingSplunk) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


